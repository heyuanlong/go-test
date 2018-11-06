package main

import (
	"time"
	"fmt"
	"github.com/nsqio/go-nsq"
)

func main() {
	go produce()
	go consume1()
	go consume2()
	go consume3()
	select {

	}
}

var topic string = "topic16"

func produce(){
	//  try to connect
	cfg := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", cfg)
	if nil != err {
		fmt.Println(err)
		return
	}

	//  try to ping
	err = producer.Ping()
	if nil != err {
		producer.Stop()
		producer = nil
		fmt.Println(err)
		return
	}
	for{
		time.Sleep(time.Second)
		err =producer.Publish(topic,[]byte("哈哈哈哈哈"))
		fmt.Println("---------------------------------------------")
		if nil != err {
			producer.Stop()
			producer = nil
			fmt.Println(err)
			return
		}
	}
}
//---------------------------------------------------------
type NSQHandler struct {
	 name string
}

func (this *NSQHandler) HandleMessage(message *nsq.Message) error {
	fmt.Println(this.name," recv:", string(message.Body))
	return nil
}
func consume1(){
	consumer,_ := nsq.NewConsumer(topic,"ch1",nsq.NewConfig())
	consumer.AddHandler(&NSQHandler{name:"consume1"})
	consumer.ConnectToNSQLookupd("127.0.0.1:4161")

	if consumer.Stats().Connections == 0{
		for{
			time.Sleep(time.Second)
			fmt.Println("try connect")
			consumer,_ = nsq.NewConsumer(topic,"ch1",nsq.NewConfig())
			consumer.AddHandler(&NSQHandler{name:"consume1"})
			consumer.ConnectToNSQLookupd("127.0.0.1:4161")

			if consumer.Stats().Connections == 1 {
				break
			}
		}
	}

	select {

	}
}
func consume2(){
	consumer,_ := nsq.NewConsumer(topic,"ch1",nsq.NewConfig())
	consumer.AddHandler(&NSQHandler{name:"consume2"})
	consumer.ConnectToNSQLookupd("127.0.0.1:4161")

	if consumer.Stats().Connections == 0{
		for{
			time.Sleep(time.Second)
			fmt.Println("try connect")
			consumer,_ = nsq.NewConsumer(topic,"ch1",nsq.NewConfig())
			consumer.AddHandler(&NSQHandler{name:"consume2"})
			consumer.ConnectToNSQLookupd("127.0.0.1:4161")

			if consumer.Stats().Connections == 1 {
				break
			}
		}
	}

	select {

	}
}
func consume3(){
	consumer,_ := nsq.NewConsumer(topic,"ch3",nsq.NewConfig())
	consumer.AddHandler(&NSQHandler{name:"consume3"})
	//consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	consumer.ConnectToNSQD("127.0.0.1:4150")

	if consumer.Stats().Connections == 0{
		for{
			time.Sleep(time.Second)
			fmt.Println("try connect")
			consumer,_ = nsq.NewConsumer(topic,"ch3",nsq.NewConfig())
			consumer.AddHandler(&NSQHandler{name:"consume3"})
			//consumer.ConnectToNSQLookupd("127.0.0.1:4161")
			consumer.ConnectToNSQD("127.0.0.1:4150")

			if consumer.Stats().Connections == 1 {
				break
			}
		}
	}
	select{

	}
}