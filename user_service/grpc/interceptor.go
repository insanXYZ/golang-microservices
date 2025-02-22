package main

import (
	"context"
	"fmt"
	"slices"

	userpb "github.com/insanXYZ/proto/gen/go/user"
	"google.golang.org/grpc"
)

func (u *UserServer) VerifyJwtInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	fmt.Println("[INTERCEPTOR USER GRPC] verify jwt")

	excludeMethods := []string{
		userpb.UserService_FindUserByEmail_FullMethodName,
		userpb.UserService_Insert_FullMethodName,
	}

	if slices.Contains(excludeMethods, info.FullMethod) {
		return handler(ctx, req)
	}

	return u.authClient.Verify(ctx, nil)
}
