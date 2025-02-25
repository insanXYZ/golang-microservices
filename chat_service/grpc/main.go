package main

import (
	"chat-service-grpc/dial"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	APP_PORT   = ":8084"
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
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(chatServer.VerifyJwtInterceptor))

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
