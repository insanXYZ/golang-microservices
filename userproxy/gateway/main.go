package main

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/insanXYZ/proto/gen/gateway/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

const (
	APP_PORT      = ":3125"
	GRPC_ENDPOINT = "localhost:3124"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := user.RegisterUsersHandlerFromEndpoint(ctx, mux, GRPC_ENDPOINT, opts)
	if err != nil {
		return err
	} // Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(APP_PORT, mux)
}

func main() {
	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
