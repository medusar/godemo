package main

import (
	"context"
	"godemo/grpc/helloworld"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

func main() {
	address := "127.0.0.1:5555"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := helloworld.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := "jason"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err = c.SayHi(ctx, &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
