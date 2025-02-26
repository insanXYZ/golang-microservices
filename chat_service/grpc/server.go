package main

import (
	"context"

	authpb "github.com/insanXYZ/proto/gen/go/auth"
	chatpb "github.com/insanXYZ/proto/gen/go/chat"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatService struct {
	authClient authpb.AuthServiceClient
	chatpb.UnimplementedChatServiceServer
}

func NewChatServer(authClient authpb.AuthServiceClient) *ChatService {
	return &ChatService{
		authClient: authClient,
	}
}

func (c *ChatService) Upgrade(_ *emptypb.Empty, stream grpc.ServerStreamingServer[chatpb.Message]) error {
	return nil
}

func (c *ChatService) BroadcastMessage(ctx context.Context, msg *chatpb.Message) (*emptypb.Empty, error) {
	return nil, nil
}
