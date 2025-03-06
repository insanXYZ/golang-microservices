module github.com/insanXYZ/chat-service-gateway

go 1.24.0

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.25.1
	github.com/insanXYZ/proto v1.0.0
	google.golang.org/grpc v1.71.0
)

require (
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250106144421-5f5ef82da422 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250303144028-a0af3efb3deb // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)

replace github.com/insanXYZ/proto => ../../proto
