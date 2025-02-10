package dial

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var opts = grpc.WithTransportCredentials(insecure.NewCredentials())
