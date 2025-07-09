package main

import (
	"context"
	"log"
	"net"

	pb "protobuf" // This should match your go_package or go.mod name

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

// Implements the SayHello RPC method
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received request for name: %s", req.GetName())
	return &pb.HelloResponse{
		Data: "hello microservices",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})

	log.Println("✅ gRPC server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
