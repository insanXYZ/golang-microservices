package main

import (
	"chat-service-grpc/dial"
	"log"
	"net"
	"os"

	chatpb "github.com/insanXYZ/proto/gen/go/chat"
	"google.golang.org/grpc"
)

var (
	APP_PORT   = os.Getenv("APP_PORT")
	LOG_PREFIX = "[CHAT GRPC]"
)

func main() {
	// config

	authClient, err := dial.NewAuthServiceClient()
	if err != nil {
		log.Fatal(LOG_PREFIX, "Error dial:", err.Error())
	}

	// server init
	chatServer := NewChatServer(authClient)
	grpcServer := grpc.NewServer(grpc.ChainStreamInterceptor(chatServer.StreamVerifyJwtInterceptor), grpc.ChainUnaryInterceptor(chatServer.UnaryVerifyJwtInterceptor))

	chatpb.RegisterChatServiceServer(grpcServer, chatServer)

	listen, err := net.Listen("tcp", APP_PORT)
	if err != nil {
		log.Fatal(LOG_PREFIX, "Error listen port", err.Error())
	}

	log.Println(LOG_PREFIX, "running on port", APP_PORT)
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal(LOG_PREFIX, "Error listen grpc:", err.Error())
	}

}
