package dial

import (
	"os"

	authpb "github.com/insanXYZ/proto/gen/go/auth"
	"google.golang.org/grpc"
)

var AUTH_SERVICE_PORT = os.Getenv("DIAL_AUTH_ENDPOINT")

func NewAuthServiceClient() authpb.AuthServiceClient {
	clientConn, err := grpc.NewClient(AUTH_SERVICE_PORT, opts)
	if err != nil {
		panic(err.Error())
	}

	return authpb.NewAuthServiceClient(clientConn)
}
