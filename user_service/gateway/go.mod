module github.com/insanXYZ/user-service-gateway

go 1.23.4

require (
	github.com/insanXYZ/proto v1.0.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.25.1
	google.golang.org/grpc v1.69.2
)

require (
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241219192143-6b3ec007d9bb // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241219192143-6b3ec007d9bb // indirect
	google.golang.org/protobuf v1.36.1 // indirect
)

replace github.com/insanXYZ/proto => ../../proto/
