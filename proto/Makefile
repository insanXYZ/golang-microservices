.PHONY: protoc
protoc:
	protoc -I . --grpc-gateway_out ./gen/gateway \
	--openapiv2_out ./gen/gateway \
	--openapiv2_opt generate_unbound_methods=true \
	--grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	--go_out ./gen/go --go-grpc_out ./gen/go \
	proto/user/*.proto proto/auth/*.proto

.PHONY: injectTag
injectTag:
	protoc-go-inject-tag -input="gen/go/*/*.pb.go"

.PHONY: generate
generate: protoc injectTag