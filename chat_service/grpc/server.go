package main

import (
	"context"

	authpb "github.com/insanXYZ/proto/gen/go/auth"
	chatpb "github.com/insanXYZ/proto/gen/go/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatService struct {
	Hub        *Hub
	authClient authpb.AuthServiceClient
	chatpb.UnimplementedChatServiceServer
}

func NewChatServer(authClient authpb.AuthServiceClient) *ChatService {
	return &ChatService{
		Hub:        NewHub(),
		authClient: authClient,
	}
}

func (c *ChatService) BroadcastMessage(ctx context.Context, msg *chatpb.Message) (*emptypb.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}
	c.Hub.Append(md["username"][0] , &Client{
		stream: ,
	})
}
