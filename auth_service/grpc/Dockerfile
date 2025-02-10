FROM golang:1.23-alpine as builder

WORKDIR /app
COPY auth_service/grpc/ /app
COPY proto /app/proto

RUN go mod edit -replace=github.com/insanXYZ/proto=/app/proto
RUN go get
RUN go build -o /app/main .

FROM alpine:3
WORKDIR /app
COPY --from=builder /app/main /app
RUN chmod +x main
CMD [ "./main" ]
