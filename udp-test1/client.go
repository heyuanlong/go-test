package main

import (
	"os"
	"fmt"
	"net"
	//  "io"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:11110")
	defer conn.Close()
	if err != nil {
		os.Exit(1)
	}

	n,err :=conn.Write([]byte("Hello world!"))
	if err != nil {
		fmt.Println("Write fail")
		os.Exit(1)
	}
	fmt.Println("send n:" ,n)

	fmt.Println("send msg")

	var msg [20]byte
	conn.Read(msg[0:])

	fmt.Println("msg is", string(msg[0:10]))
}