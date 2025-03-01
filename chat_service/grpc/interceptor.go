package main

import (
	"fmt"

	chatpb "github.com/insanXYZ/proto/gen/go/chat"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (c *ChatService) VerifyJwtInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	fmt.Println(LOG_PREFIX, "verify jwt interceptor")

	if info.FullMethod == chatpb.ChatService_BroadcastMessage_FullMethodName {
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

	return status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
}
