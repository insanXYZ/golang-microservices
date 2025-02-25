package main

import (
	"context"
	"log"
	"net"
	"os"

	userpb "github.com/insanXYZ/proto/gen/go/user"
	"github.com/insanXYZ/user-service-grpc/config"
	"github.com/insanXYZ/user-service-grpc/dial"
	"google.golang.org/grpc"
)

var (
	APP_PORT   = os.Getenv("APP_PORT")
	LOG_PREFIX = "[GRPC USER]"
)

func main() {
	ctx := context.Background()

	// config
	pgxConn, err := config.NewDatabase(ctx)
	if err != nil {
		log.Fatal(LOG_PREFIX, "Error connect database :", err.Error())
	}
	validator := config.NewValidator()

	// dial client
	authClient, err := dial.NewAuthServiceClient()
	if err != nil {
		log.Fatal(LOG_PREFIX, "Error dial :", err.Error())
	}
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
		log.Fatal(LOG_PREFIX, "Error listen port :", err.Error())
	}

	log.Println(LOG_PREFIX, "listen on port "+APP_PORT)
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal(LOG_PREFIX, "Error server grpc server :", err.Error())
	}
}
