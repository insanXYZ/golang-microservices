package main

import (
	"context"
	"slices"

	userpb "github.com/insanXYZ/proto/gen/go/user"
	"google.golang.org/grpc"
)

func (u *UserServer) VerifyJwtInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	LogPrintln("verify jwt interceptor")

	excludeMethods := []string{
		userpb.UserService_FindUserByEmail_FullMethodName,
		userpb.UserService_Insert_FullMethodName,
	}

	if slices.Contains(excludeMethods, info.FullMethod) {
		return handler(ctx, req)
	}

	_, err = u.authClient.Verify(ctx, nil)

	if err != nil {
		LogPrintln("Error verify auth", err.Error())
		return nil, err
	}

	return handler(ctx, req)
}
