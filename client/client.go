package main

import (
	"context"
	"log"
	"os"

	pb "github.com/knry0329/grpc-hello-world"
	"google.golang.org/grpc"
)

func main() {
	addr := "localhost:50051"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %c", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := os.Args[1]
	age := int32(123)

	ctx := context.Background()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name, Age: age})
	if err != nil {
		log.Fatalf("could not greet: %c", err)

	}
	log.Printf("Greeting: %s", r.Message)
}
