package main

import (
	"context"
	"log"
	"net"

	usersv "github.com/insanXYZ/proto/gen/go/user"
	"google.golang.org/grpc"
)

const APP_PORT = ":3124"

type UserServer struct {
	usersv.UnimplementedUsersServiceServer
}

func (u *UserServer) FindUserByEmail(context.Context, *usersv.FindUserByEmailRequest) (*usersv.FindUserByEmailResponse, error) {

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
