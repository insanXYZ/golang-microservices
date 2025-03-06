FROM golang:1.24-alpine as builder

WORKDIR /app
COPY user_service/gateway/ /app
COPY proto /app/proto

RUN go mod edit -replace=github.com/insanXYZ/proto=/app/proto
RUN go get
RUN go build -o /app/main .

FROM alpine:3
WORKDIR /app
COPY --from=builder /app/main /app
RUN chmod +x main
CMD [ "./main" ]
