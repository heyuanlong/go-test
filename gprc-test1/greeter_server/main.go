
package main

import (
	"log"
	"net"
	"fmt"

	"google.golang.org/grpc"
	pb "go-test/gprc-test1/helloworld"
	"google.golang.org/grpc/reflection"
	"time"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func (s *server) SayHelloClientToServer(r pb.Greeter_SayHelloClientToServerServer) error{
	for  {
		req , err := r.Recv()
		if err != nil{
			fmt.Println(err)
			return err
		}
		fmt.Println(req)
	}
	return nil
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHelloServerToClient(in *pb.HelloRequest, gserver pb.Greeter_SayHelloServerToClientServer) error {
	fmt.Println(in)
	resp := &pb.HelloReply{Message: "ServerToClient:sTOc "}
	for  {
		time.Sleep(time.Second * 1)
		gserver.Send(resp)
	}
	return nil
}

func (s *server) SayHelloServerToServer(ss pb.Greeter_SayHelloServerToServerServer) error{
	resp := &pb.HelloReply{Message: "ServerToServer:sTOc"}
	for  {
		req ,err := ss.Recv()
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(req)
		err = ss.Send(resp)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
