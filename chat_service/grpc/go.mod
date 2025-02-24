module chat-service-grpc

go 1.24.0

require github.com/insanXYZ/proto v1.0.0

require (
	github.com/gorilla/websocket v1.5.3 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250219182151-9fdb1cabc7b2 // indirect
	google.golang.org/grpc v1.70.0 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)

replace github.com/insanXYZ/proto => ../../proto/
