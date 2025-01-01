package main

import (
	"context"
	"log"
	"net"

	usersv "github.com/insanXYZ/proto/gen/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const APP_PORT = ":3124"

type UserServer struct {
	usersv.UnimplementedUsersServer
}

func (u *UserServer) Login(_ context.Context, req *usersv.UserLoginRequest) (*usersv.UserLoginResponse, error) {
	if req.Username == "insan" && req.Password == "insan1601" {
		return &usersv.UserLoginResponse{
			Status: "success login",
		}, nil
	}

	return nil, status.Error(codes.Unimplemented, "error login")
}

func main() {
	server := grpc.NewServer()
	var userServer UserServer
	usersv.RegisterUsersServer(server, &userServer)

	listen, err := net.Listen("tcp", APP_PORT)
	if err != nil {
		panic(err.Error())
	}

	log.Printf("grpc userserver listen on port %s", APP_PORT)

	err = server.Serve(listen)
	if err != nil {
		panic(err.Error())
	}
}
