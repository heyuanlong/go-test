package main

import (
	"net"
	"time"
	"fmt"
)


func main() {
	fmt.Println("start")
	listen_sock, err := net.Listen("tcp",":"+"1234")
	if err != nil{
		fmt.Println(err)
	}
	defer listen_sock.Close()
	for{
		new_conn, err := listen_sock.Accept()
		if err != nil {
			fmt.Println("listen_sock.Accept error:", err)
			continue
		}
		go HandleClient(new_conn)
	}
}

func HandleClient(conn net.Conn)  {
	fmt.Println("HandleClient")
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("HandleClient func s")
		conn.Close()
		fmt.Println("HandleClient func e")
	}()

	defer func() {
		conn.Close()
		fmt.Println("HandleClient defer")
	}()

	var msgBuf = make([]byte, 1024)
	for {
		conn.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))
		n, err := conn.Read(msgBuf)
		if err != nil {
			if nerr, ok := err.(*net.OpError); ok && nerr.Timeout() {
				fmt.Println("timeout")
				return
			} else {
				fmt.Println("read close or fail")
				return
			}
		}
		fmt.Println(n)
	}
}