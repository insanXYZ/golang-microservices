.PHONY: protoc
protoc:
	rm -rf gen/gateway/
	mkdir gen/gateway/
	protoc -I . --grpc-gateway_out ./gen/gateway \
	--openapiv2_out ./gen/gateway \
	--openapiv2_opt generate_unbound_methods=true \
	--grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	--go_out ./gen/go --go-grpc_out ./gen/go \
	proto/user/*.proto proto/auth/*.proto proto/chat/*.proto

.PHONY: injectTag
injectTag:
	protoc-go-inject-tag -input="gen/go/*/*.pb.go"

.PHONY: initGatewayDir
initGatewayDir:
	mv gen/gateway/proto/auth/* gen/go/auth
	mv gen/gateway/proto/user/* gen/go/user
	mv gen/gateway/proto/chat/* gen/go/chat
	rm -rf gen/gateway/

.PHONY: generate
generate: protoc injectTag initGatewayDir 