package main

import (
	"godemo/grpc/helloworld"
	"godemo/grpc/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := ":5555"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server.GreetServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
