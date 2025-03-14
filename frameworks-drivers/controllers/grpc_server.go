package controllers

import (
	"context"
	"fmt"
	interface_adapters "go-boilerplate-clean-architecture/interface-adapters"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	interface_adapters.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *interface_adapters.HelloRequest) (*interface_adapters.HelloResponse, error) {
	name := req.GetName()
	message := fmt.Sprintf("Hello, %s!", name)
	return &interface_adapters.HelloResponse{Message: message}, nil
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("gRPC failed to listen: %v", err)
	}

	s := grpc.NewServer()
	interface_adapters.RegisterHelloServiceServer(s, &server{})

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("gRPC failed to serve: %v", err)
	}
}
