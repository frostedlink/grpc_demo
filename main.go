package main

import (
	"context"
	"fmt"
	proto "grpc_demo_go/proto"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type chatServer struct {
	proto.UnimplementedChatServiceServer
}

func (s *chatServer) SendMessage(ctx context.Context, message *proto.Message) (*proto.Message, error) {
	log.Printf("Received message body: %s", message.Body)
	return &proto.Message{Body: "Hello from server"}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterChatServiceServer(grpcServer, &chatServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
