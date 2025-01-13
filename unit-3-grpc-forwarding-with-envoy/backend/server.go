package main

import (
	"context"
	"fmt"
	"log"
	"net"

	chatpb "backend/module/chatpb"

	"google.golang.org/grpc"
)

type server struct {
	chatpb.UnimplementedChatServiceServer
}

func (s *server) SayHello(ctx context.Context, req *chatpb.HelloRequest) (*chatpb.HelloResponse, error) {
	message := fmt.Sprintf("Hello, %s!", req.Name)
	return &chatpb.HelloResponse{Message: message}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	chatpb.RegisterChatServiceServer(grpcServer, &server{})
	log.Println("gRPC server running on port 5001...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
