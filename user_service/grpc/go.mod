module github.com/insanXYZ/user-service-grpc

go 1.23.4

require (
	github.com/go-playground/validator/v10 v10.24.0
	github.com/google/uuid v1.6.0
	github.com/insanXYZ/proto v1.0.0
	github.com/jackc/pgx/v5 v5.7.2
	golang.org/x/crypto v0.32.0
	google.golang.org/grpc v1.69.2
)

require (
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.25.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241219192143-6b3ec007d9bb // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241219192143-6b3ec007d9bb // indirect
	google.golang.org/protobuf v1.36.1 // indirect
)

replace github.com/insanXYZ/proto => ../../proto/
