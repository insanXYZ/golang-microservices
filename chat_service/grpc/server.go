package main

import (
	"log"

	authpb "github.com/insanXYZ/proto/gen/go/auth"
	chatpb "github.com/insanXYZ/proto/gen/go/chat"
	"github.com/insanXYZ/proto/gen/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ChatService struct {
	Hub        *Hub
	authClient authpb.AuthServiceClient
	chatpb.UnimplementedChatServiceServer
}

func NewChatServer(authClient authpb.AuthServiceClient) *ChatService {
	hub := NewHub()
	go hub.Run()

	return &ChatService{
		Hub:        hub,
		authClient: authClient,
	}
}

func (c *ChatService) BroadcastMessage(stream grpc.BidiStreamingServer[chatpb.MessageRequest, chatpb.MessageResponse]) error {
	log.Println("Using broadcastmessage rpc")
	ctx := stream.Context()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("Error missing getting header from incoming context")
		return status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}

	var client *Client

	if !c.Hub.ExistClient(md.Get("id")[0]) {
		client = &Client{
			stream: stream,
			user: &user.User{
				Id:   md.Get("id")[0],
				Name: md.Get("name")[0],
			},
			Hub: c.Hub,
			err: make(chan error),
		}

		client.Hub.Register <- client
	} else {
		client = c.Hub.Clients[md.Get("id")[0]]
	}

	go client.ReadPump()
	return <-client.err
}
