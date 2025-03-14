package main

import (
	"errors"
	"io"
	"log"
	"time"

	chatpb "github.com/insanXYZ/proto/gen/go/chat"
	userpb "github.com/insanXYZ/proto/gen/go/user"
	"google.golang.org/grpc"
)

type Client struct {
	stream grpc.BidiStreamingServer[chatpb.MessageRequest, chatpb.MessageResponse]
	Hub    *Hub
	user   *userpb.User
	err    chan error
}

func (c *Client) ReadPump() {
	for {
		select {
		case <-c.stream.Context().Done():
			log.Printf("Client with name %v canceled", c.user.Name)
			c.err <- errors.New("Client canceled")
			return
		default:
			msg, err := c.stream.Recv()
			if err == io.EOF {
				log.Println("Error EOF", c.user.Name)
				c.err <- errors.New("Error eof")
				break
			} else if err != nil {
				log.Println("Error recv", err.Error(), c.user.Name)
				c.err <- errors.New("Error recv")
				return
			}

			msgResponse := chatpb.MessageResponse{
				Message:   msg.Message,
				Timestamp: time.Now().Format(time.DateTime),
				User:      c.user,
			}

			c.Hub.Broadcast <- &msgResponse
		}
	}
}
