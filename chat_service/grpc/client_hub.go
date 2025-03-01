package main

import (
	chatpb "github.com/insanXYZ/proto/gen/go/chat"
	userpb "github.com/insanXYZ/proto/gen/go/user"
	"google.golang.org/grpc"
)

type Client struct {
	stream grpc.ServerStreamingServer[chatpb.Message]
	user   *userpb.User
	err    chan error
}
