package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)
import "github.com/ishmaiah-ih/"

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServerService
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start: %v", err)
	}

}
