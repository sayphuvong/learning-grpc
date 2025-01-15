package main

import (
	"context"
	"log"
	"time"

	chatpb "backend/module/chatpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Envoy: %v", err)
	}
	defer conn.Close()

	client := chatpb.NewChatServiceClient(conn)
	req := &chatpb.HelloRequest{Name: "Envoy"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}
	log.Printf("Response from server: %s", res.Message)
}
