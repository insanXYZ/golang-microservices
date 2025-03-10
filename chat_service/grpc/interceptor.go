package main

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type wrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (s *wrappedStream) Context() context.Context {
	return s.ctx
}

func (c *ChatService) StreamVerifyJwtInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	LogPrintln("using StreamVerifyJwtInterceptor")
	md, err := c.verifyJwt(ss.Context())
	if err != nil {
		LogPrintln("Error verify jwt on StreamVerifyJwtInterceptor :", err.Error())
		return err
	}

	ctx := metadata.NewIncomingContext(ss.Context(), md)

	return handler(srv, &wrappedStream{
		ctx:          ctx,
		ServerStream: ss,
	})

}

func (c *ChatService) UnaryVerifyJwtInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	LogPrintln("using UnaryVerifyJwtInterceptor")
	md, err := c.verifyJwt(ctx)
	if err != nil {
		LogPrintln("Error verify jwt on UnaryVerifyJwtInterceptor :", err.Error())
		return nil, status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}

	ctx = metadata.NewIncomingContext(ctx, md)

	return handler(ctx, req)
}

// called this function for verify jwt to auth service fw
// and get jwt claims (user.User model) for next handler
func (c *ChatService) verifyJwt(ctx context.Context) (metadata.MD, error) {
	LogPrintln("verifyJwt")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		LogPrintln("Error getting metadata on context")
		return nil, errors.New("Error getting metadata on context")
	}
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp, err := c.authClient.Verify(ctx, nil)
	if err != nil {
		LogPrintln("Error verify auth", err.Error())
		return nil, err
	}

	user := resp.User

	md, ok = metadata.FromOutgoingContext(ctx)
	if !ok {
		LogPrintln("Error from outgoingcontext on context")
		return nil, errors.New("Error getting metadata on context")
	}
	md.Append("name", user.GetName())
	md.Append("id", user.GetId())

	return md, nil
}
