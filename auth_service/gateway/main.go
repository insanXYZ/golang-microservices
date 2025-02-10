package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	authpb "github.com/insanXYZ/proto/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	APP_PORT      = os.Getenv("APP_PORT")
	GRPC_ENDPOINT = os.Getenv("GRPC_ENDPOINT")
)

func run() error {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, GRPC_ENDPOINT, opts)
	if err != nil {
		return err
	}

	fmt.Println("authproxy[GATEWAY] listen on port " + APP_PORT)
	return http.ListenAndServe(APP_PORT, mux)
}

func main() {
	if err := run(); err != nil {
		panic(err.Error())
	}
}
