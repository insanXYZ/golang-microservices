package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	chatpb "github.com/insanXYZ/proto/gen/go/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	APP_PORT      = os.Getenv("APP_PORT")
	GRPC_ENDPOINT = os.Getenv("GRPC_ENDPOINT")
	LOG_PREFIX    = "[CHAT GATEWAY]"
)

func run() error {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := chatpb.RegisterChatServiceHandlerFromEndpoint(ctx, mux, GRPC_ENDPOINT, opts)
	if err != nil {
		return err
	}

	fmt.Println(LOG_PREFIX, "running on port ", APP_PORT)
	return http.ListenAndServe(APP_PORT, mux)
}

func main() {
	if err := run(); err != nil {
		panic(err.Error())
	}
}
