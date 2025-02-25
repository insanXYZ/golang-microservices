package dial

import (
	"os"

	userpb "github.com/insanXYZ/proto/gen/go/user"
	"google.golang.org/grpc"
)

var USER_SERVICE_PORT = os.Getenv("DIAL_USER_ENDPOINT")

func NewUserServiceClient() (userpb.UserServiceClient, error) {
	clientConn, err := grpc.NewClient(USER_SERVICE_PORT, opts)
	if err != nil {
		return nil, err
	}

	return userpb.NewUserServiceClient(clientConn), nil
}
