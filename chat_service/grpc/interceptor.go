package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func (c *ChatService) VerifyJwtInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	log.Println(LOG_PREFIX, "verify jwt interceptor")
	return c.authClient.Verify(ctx, nil)
}
