package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	authsv "github.com/insanXYZ/proto/gen/gateway/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	APP_PORT      = ":8080"
	GRPC_ENDPOINT = "auth_grpc:8081"
)

func run() error {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := authsv.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, GRPC_ENDPOINT, opts)
	if err != nil {
		return err
	}

	fmt.Println("authproxy[GATEWAY] listen on port " + APP_PORT)
	err = http.ListenAndServe(APP_PORT, mux)
	if err != nil {
		return err
	}
}

func main() {
	if err := run(); err != nil {
		panic(err.Error())
	}
}
