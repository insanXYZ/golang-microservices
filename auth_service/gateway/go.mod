module github.com/insanXYZ/auth-service-gateway

go 1.23.5

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.25.1
	google.golang.org/grpc v1.70.0
	github.com/insanXYZ/proto v1.0.0
)

require (
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241219192143-6b3ec007d9bb // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250207221924-e9438ea467c6 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)

replace github.com/insanXYZ/proto => ../../proto
