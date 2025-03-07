package main

import (
	"auth-service-grpc/dial"
	"log"
	"net"

	"github.com/go-playground/validator/v10"
	authsv "github.com/insanXYZ/proto/gen/go/auth"
	"google.golang.org/grpc"
)

const (
	APP_PORT   = ":8081"
	LOG_PREFIX = "[GRPC AUTH]"
)

func main() {
	grpcServer := grpc.NewServer()

	userClient, err := dial.NewUserServiceClient()
	if err != nil {
		log.Fatal(LOG_PREFIX, "Error dial :", err.Error())
	}

	chatClient, err := dial.NewChatServiceClient()
	if err != nil {
		log.Fatal(LOG_PREFIX, "Error dial :", err.Error())
	}

	initServer := NewAuthServer(userClient, chatClient, validator.New())
	authsv.RegisterAuthServiceServer(grpcServer, initServer)

	listen, err := net.Listen("tcp", APP_PORT)
	if err != nil {
		log.Fatal(LOG_PREFIX, "Error listen port :", err.Error())
	}

	log.Println(LOG_PREFIX, "listen on port "+APP_PORT)
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal(LOG_PREFIX, "Error server grpc :", err.Error())
	}

}
