package dial

import (
	usersv "github.com/insanXYZ/proto/gen/go/user"
	"google.golang.org/grpc"
)

const USER_SERVICE_PORT = "localhost:8083"

func NewUserServiceClient() usersv.UserServiceClient {
	clientConn, err := grpc.NewClient(USER_SERVICE_PORT, opts)
	if err != nil {
		panic(err.Error())
	}

	return usersv.NewUserServiceClient(clientConn)
}
