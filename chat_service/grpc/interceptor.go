package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (c *ChatService) StreamVerifyJwtInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	fmt.Println(LOG_PREFIX, "verify jwt interceptor")

	resp, err := c.authClient.Verify(ss.Context(), nil)
	if err != nil {
		return status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}

	user := resp.User
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}
	md.Append("username", user.GetUsername())
	md.Append("email", user.GetEmail())
	md.Append("id", user.GetId())
	err = ss.SetHeader(md)
	if err != nil {
		return status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}

	return handler(srv, ss)

}

func (c *ChatService) UnaryVerifyJwtInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

}

// called this function for verify jwt to auth service
// and get jwt claims (user.User model) for next handler
func (c *ChatService) verifyJwt(ctx context.Context) (metadata.MD, error) {
	resp, err := c.authClient.Verify(ctx, nil)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
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
