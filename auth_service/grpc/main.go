package main

import (
	"auth-service-grpc/dial"
	"fmt"
	"net"

	"github.com/go-playground/validator/v10"
	authsv "github.com/insanXYZ/proto/gen/go/auth"
	"google.golang.org/grpc"
)

const (
	APP_PORT = ":8081"
)

func main() {
	grpcServer := grpc.NewServer()

	initServer := NewAuthServer(dial.NewUserServiceClient(), validator.New())

	authsv.RegisterAuthServiceServer(grpcServer, initServer)

	listen, err := net.Listen("tcp", APP_PORT)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("authproxy[GRPC] listen on port " + APP_PORT)
	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err.Error())
	}

}
