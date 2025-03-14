// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: proto/chat/chat.proto

/*
Package chat is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package chat

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var (
	_ codes.Code
	_ io.Reader
	_ status.Status
	_ = errors.New
	_ = runtime.String
	_ = utilities.NewDoubleArray
	_ = metadata.Join
)

func request_ChatService_BroadcastMessage_0(ctx context.Context, marshaler runtime.Marshaler, client ChatServiceClient, req *http.Request, pathParams map[string]string) (ChatService_BroadcastMessageClient, runtime.ServerMetadata, chan error, error) {
	var metadata runtime.ServerMetadata
	errChan := make(chan error, 1)
	stream, err := client.BroadcastMessage(ctx)
	if err != nil {
		grpclog.Errorf("Failed to start streaming: %v", err)
		close(errChan)
		return nil, metadata, errChan, err
	}
	dec := marshaler.NewDecoder(req.Body)
	handleSend := func() error {
		var protoReq MessageRequest
		err := dec.Decode(&protoReq)
		if errors.Is(err, io.EOF) {
			return err
		}
		if err != nil {
			grpclog.Errorf("Failed to decode request: %v", err)
			return status.Errorf(codes.InvalidArgument, "Failed to decode request: %v", err)
		}
		if err := stream.Send(&protoReq); err != nil {
			grpclog.Errorf("Failed to send request: %v", err)
			return err
		}
		return nil
	}
	go func() {
		defer close(errChan)
		for {
			if err := handleSend(); err != nil {
				errChan <- err
				break
			}
		}
		if err := stream.CloseSend(); err != nil {
			grpclog.Errorf("Failed to terminate client stream: %v", err)
		}
	}()
	header, err := stream.Header()
	if err != nil {
		grpclog.Errorf("Failed to get header from client: %v", err)
		return nil, metadata, errChan, err
	}
	metadata.HeaderMD = header
	return stream, metadata, errChan, nil
}

// RegisterChatServiceHandlerServer registers the http handlers for service ChatService to "mux".
// UnaryRPC     :call ChatServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterChatServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterChatServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server ChatServiceServer) error {
	mux.Handle(http.MethodPost, pattern_ChatService_BroadcastMessage_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterChatServiceHandlerFromEndpoint is same as RegisterChatServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterChatServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterChatServiceHandler(ctx, mux, conn)
}

// RegisterChatServiceHandler registers the http handlers for service ChatService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterChatServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterChatServiceHandlerClient(ctx, mux, NewChatServiceClient(conn))
}

// RegisterChatServiceHandlerClient registers the http handlers for service ChatService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ChatServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ChatServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ChatServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterChatServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ChatServiceClient) error {
	mux.Handle(http.MethodPost, pattern_ChatService_BroadcastMessage_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/chat.ChatService/BroadcastMessage", runtime.WithHTTPPathPattern("/api/chat/broadcast"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		resp, md, reqErrChan, err := request_ChatService_BroadcastMessage_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		go func() {
			for err := range reqErrChan {
				if err != nil && !errors.Is(err, io.EOF) {
					runtime.HTTPStreamError(annotatedContext, mux, outboundMarshaler, w, req, err)
				}
			}
		}()
		forward_ChatService_BroadcastMessage_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_ChatService_BroadcastMessage_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"api", "chat", "broadcast"}, ""))
)

var (
	forward_ChatService_BroadcastMessage_0 = runtime.ForwardResponseStream
)
