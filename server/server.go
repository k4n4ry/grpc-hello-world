package main

import (
	"context"
	"log"
	"net"
	"strconv"

	pb "github.com/knry0329/grpc-hello-world"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Recieved: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name + ", your age is " + strconv.Itoa(int(in.Age))}, nil

}

func main() {
	addr := ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("gRPC server listening on " + addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
