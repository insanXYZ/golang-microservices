package main

import (
	"context"
	"fmt"
	"net"
	"os"

	userpb "github.com/insanXYZ/proto/gen/go/user"
	"github.com/insanXYZ/user-service-grpc/config"
	"github.com/insanXYZ/user-service-grpc/dial"
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

	// dial client
	authClient := dial.NewAuthServiceClient()

	// main server
	userServer := NewUserServer(pgxConn, authClient, validator)
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			userServer.VerifyJwtInterceptor,
		),
	)

	// register server
	userpb.RegisterUserServiceServer(grpcServer, userServer)

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
