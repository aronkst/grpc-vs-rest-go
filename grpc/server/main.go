package main

import (
	"log"
	"net"

	"github.com/aronkst/grpc-vs-rest-go/grpc/message"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	message.RegisterMessageServiceServer(grpcServer, &message.Server{})

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
