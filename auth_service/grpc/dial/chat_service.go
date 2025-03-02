package dial

import (
	"os"

	chatpb "github.com/insanXYZ/proto/gen/go/chat"
	"google.golang.org/grpc"
)

var CHAT_SERVICE_PORT = os.Getenv("DIAL_CHAT_ENDPOINT")

func NewChatServiceClient() (chatpb.ChatServiceClient, error) {
	clientConn, err := grpc.NewClient(USER_SERVICE_PORT, opts)
	if err != nil {
		return nil, err
	}

	return chatpb.NewChatServiceClient(clientConn), nil
}
