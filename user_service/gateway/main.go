package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	usersv "github.com/insanXYZ/proto/gen/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

var (
	APP_PORT      = os.Getenv("APP_PORT")
	GRPC_ENDPOINT = os.Getenv("GRPC_ENDPOINT")
)

func run() error {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := usersv.RegisterUserServiceHandlerFromEndpoint(ctx, mux, GRPC_ENDPOINT, opts)
	if err != nil {
		return err
	}

	fmt.Println("userproxy[GATEWAY] listen on port " + APP_PORT)
	return http.ListenAndServe(APP_PORT, mux)
}

func main() {
	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
