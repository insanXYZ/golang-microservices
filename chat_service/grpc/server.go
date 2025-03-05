package main

import (
	"context"
	"log"
	"sync"

	authpb "github.com/insanXYZ/proto/gen/go/auth"
	chatpb "github.com/insanXYZ/proto/gen/go/chat"
	userpb "github.com/insanXYZ/proto/gen/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatService struct {
	Hub        Hub
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
	wg := sync.WaitGroup{}

	for _, v := range c.Hub {
		wg.Add(1)
		go func(client *Client, msg *chatpb.Message) {
			defer wg.Done()
			err := client.stream.Send(msg)
			if err != nil {
				log.Println(LOG_PREFIX, "Error send message", err.Error())
				delete(c.Hub, v.user.Id)
				client.err <- err
			}
		}(v, msg)
	}

	wg.Wait()

	return nil, nil
}

func (c *ChatService) Subscribe(_ *emptypb.Empty, stream grpc.ServerStreamingServer[chatpb.Message]) error {
	ctx := stream.Context()
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		log.Println(LOG_PREFIX, "Error reading metadata on Subscribe rpc")
		return status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}

	client := &Client{
		stream: stream,
		user: &userpb.User{
			Id:       md.Get("id")[0],
			Username: md.Get("username")[0],
			Email:    md.Get("email")[0],
		},
		err: make(chan error),
	}

	c.Hub.Append(md.Get("id")[0], client)

	return <-client.err

}
