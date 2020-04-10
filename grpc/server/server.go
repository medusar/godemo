package server

import (
	"context"
	"godemo/grpc/helloworld"
	"log"
)

type GreetServer struct {
	helloworld.UnimplementedGreeterServer
}

func (s *GreetServer) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Println("SayHello called, req:", req.GetName())
	return &helloworld.HelloReply{Message: "Hello, " + req.GetName(),}, nil
}

func (s *GreetServer) SayHi(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Println("SayHi called, req:", req.GetName())
	return &helloworld.HelloReply{Message: "Hi, " + req.GetName(),}, nil
}
