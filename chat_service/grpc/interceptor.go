package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (c *ChatService) StreamVerifyJwtInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	md, err := c.verifyJwt(ss.Context())
	if err != nil {
		log.Println(LOG_PREFIX, "Error verify jwt on StreamVerifyJwtInterceptor :", err.Error())
		return err
	}

	err = ss.SetHeader(md)
	if err != nil {
		log.Println(LOG_PREFIX, "Error setheader metadata on StreamVerifyJwtInterceptor")
		return status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}

	return handler(srv, ss)

}

func (c *ChatService) UnaryVerifyJwtInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	md, err := c.verifyJwt(ctx)
	if err != nil {
		log.Println(LOG_PREFIX, "Error verify jwt on UnaryVerifyJwtInterceptor :", err.Error())
		return nil, err
	}

	err = grpc.SetHeader(ctx, md)
	if err != nil {
		log.Println(LOG_PREFIX, "Error setheader metadata on UnaryVerifyJwtInterceptor")
		return nil, status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}

	return handler(ctx, req)
}

// called this function for verify jwt to auth service fw
// and get jwt claims (user.User model) for next handler
func (c *ChatService) verifyJwt(ctx context.Context) (metadata.MD, error) {
	fmt.Println(LOG_PREFIX, "verifyJwt")
	resp, err := c.authClient.Verify(ctx, nil)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	user := resp.User
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}
	md.Append("username", user.GetUsername())
	md.Append("email", user.GetEmail())
	md.Append("id", user.GetId())

	return md, nil
}
