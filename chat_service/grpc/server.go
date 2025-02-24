package main

import (
	"context"

	authpb "github.com/insanXYZ/proto/gen/go/auth"
	chatpb "github.com/insanXYZ/proto/gen/go/chat"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatService struct {
	authClient authpb.AuthServiceClient
	chatpb.UnimplementedChatServiceServer
}

func NewChatService(authClient authpb.AuthServiceClient) chatpb.ChatServiceServer {
	return &ChatService{
		authClient: authClient,
	}
}

func (c *ChatService) Upgrade(ctx context.Context, _ *emptypb.Empty) (*chatpb.UpgradeResponse, error) {
	return nil, nil
}
