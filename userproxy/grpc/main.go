package main

import (
	"fmt"
	"net"

	usersv "github.com/insanXYZ/proto/gen/go/user"
	"google.golang.org/grpc"
)

const APP_PORT = ":8083"

func main() {
	grpcServer := grpc.NewServer()

	usersv.RegisterUserServiceServer(grpcServer, NewUserServer())

	listen, err := net.Listen("tcp", APP_PORT)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("userproxy[GRPC] listen on port " + APP_PORT)
	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err.Error())
	}
}
