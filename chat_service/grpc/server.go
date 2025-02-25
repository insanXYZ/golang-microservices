package main

import (
	authpb "github.com/insanXYZ/proto/gen/go/auth"
	chatpb "github.com/insanXYZ/proto/gen/go/chat"
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
