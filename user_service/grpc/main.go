package main

import (
	"context"
	"fmt"
	"net"
	"os"

	usersv "github.com/insanXYZ/proto/gen/go/user"
	"github.com/insanXYZ/user-service-grpc/config"
	"google.golang.org/grpc"
)

var (
	APP_PORT = os.Getenv("APP_PORT")
)

func main() {
	ctx := context.Background()

	// config
	pgxConn, err := config.NewDatabase(ctx)
	if err != nil {
		panic(err.Error())
	}
	validator := config.NewValidator()

	// main server
	grpcServer := grpc.NewServer()
	userServer := NewUserServer(pgxConn, validator)

	// register server
	usersv.RegisterUserServiceServer(grpcServer, userServer)

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
