package main

import (
	"context"
	"log"
	"net"

	authsv "github.com/insanXYZ/proto/gen/go/auth"
	"google.golang.org/grpc"
)

const PORT = ":8081"

type AuthServer struct {
	authsv.UnimplementedAuthServiceServer
}

func (s *AuthServer) Register(context.Context, *authsv.RegisterRequest) (*authsv.RegisterResponse, error) {
	return &authsv.RegisterResponse{
		Status: "success to Register rpc service",
	}, nil
}

func (s *AuthServer) Login(context.Context, *authsv.LoginRequest) (*authsv.LoginResponse, error) {
	return &authsv.LoginResponse{
		Status: "success to Login rpc service",
	}, nil
}

func main() {
	server := grpc.NewServer()
	authsv.RegisterAuthServiceServer(server, &AuthServer{})

	listen, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err.Error())
	}

	log.Printf("AuthService[GRPC] Run on port %s", PORT)

	err = server.Serve(listen)
	if err != nil {
		panic(err.Error())
	}

}
