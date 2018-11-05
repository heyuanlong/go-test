package  main

import (
	"context"
	"time"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:2378", "localhost:2377"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	fmt.Println("connect succ")
	defer cli.Close()

	//---------------------------------------------------- put get
	_, err = cli.Put(context.TODO(), "foo", "bar")
	if err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := cli.Get(ctx, "foo")
	cancel()
	if err != nil {
		fmt.Println(err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}

	//----------------------------------------------------  distributed locks
	go func() {
		s1, err := concurrency.NewSession(cli)
		if err != nil {
			fmt.Print(err)
		}
		defer s1.Close()
		m1 := concurrency.NewMutex(s1, "/my-lock/")


		s2, err := concurrency.NewSession(cli)
		if err != nil {
			fmt.Print(err)
		}
		defer s2.Close()
		m2 := concurrency.NewMutex(s2, "/my-lock/")


		// acquire lock for s1
		if err := m1.Lock(context.TODO()); err != nil {
			fmt.Print(err)
		}
		fmt.Println("acquired lock for s1")

		m2Locked := make(chan struct{})
		go func() {
			defer close(m2Locked)
			// wait until s1 is locks /my-lock/
			if err := m2.Lock(context.TODO()); err != nil {
				fmt.Print(err)
			}
		}()

		time.Sleep(2 * time.Second)
		if err := m1.Unlock(context.TODO()); err != nil {
			fmt.Print(err)
		}
		fmt.Println("released lock for s1")

		<-m2Locked
		fmt.Println("acquired lock for s2")
		// etcdctl --endpoints=$ENDPOINTS lock /my-lock/
		time.Sleep(5*time.Second)
		fmt.Println("s2 lock out")
	}()

	//---------------------------------------------------- work
	ctx = context.TODO()
	ch := cli.Watch(ctx, "foo", clientv3.WithPrefix())
	for {
		fmt.Print("rev\n")
		select {
		case c := <-ch:
			for _, ev := range c.Events {
				fmt.Printf("%+v\n", ev)
				fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}
}