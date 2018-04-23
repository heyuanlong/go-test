
package main

import (
	"log"
	"os"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "go-test/gprc-test1/helloworld"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		defer cancel()
		r, _ := c.SayHelloServerToClient(ctx, &pb.HelloRequest{Name: name})
		for  {
			resp , err := r.Recv()
			if err != nil{
				fmt.Println(err)
				return
			}
			fmt.Println(resp)
		}
	}()

	go func() {
		req := &pb.HelloRequest{Name: "ClientToServer:cTOs" }
		ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
		defer cancel()
		c, _ := c.SayHelloClientToServer(ctx)
		for  {
			time.Sleep(time.Second * 1)
			err := c.Send(req)
			if err != nil{
				fmt.Println(err)
				return
			}
		}
	}()
	go func() {
		req := &pb.HelloRequest{Name: "ServerToServer:cTOs" }
		ctx, cancel := context.WithTimeout(context.Background(),  20 * time.Second)
		defer cancel()
		ss, _ := c.SayHelloServerToServer(ctx)
		for  {
			err := ss.Send(req)
			if err != nil{
				fmt.Println(err)
				return
			}
			resp , err := ss.Recv()
			if err != nil{
				fmt.Println(err)
				return
			}
			fmt.Println(resp)
			time.Sleep(time.Second * 1)
		}
	}()



	time.Sleep(time.Hour)
}
