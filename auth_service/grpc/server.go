package main

import (
	"context"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	authpb "github.com/insanXYZ/proto/gen/go/auth"
	userpb "github.com/insanXYZ/proto/gen/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthServer struct {
	userClient userpb.UserServiceClient
	validator  *validator.Validate
	authpb.UnimplementedAuthServiceServer
}

func NewAuthServer(userClient userpb.UserServiceClient, validator *validator.Validate) *AuthServer {
	return &AuthServer{
		userClient: userClient,
		validator:  validator,
	}
}

func (s *AuthServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	err := s.validator.Struct(req)
	if err != nil {
		return nil, err
	}

	res, err := s.userClient.Insert(ctx, &userpb.InsertRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})

	if err != nil {
		return nil, err
	}

	return &authpb.RegisterResponse{
		Message: res.Message,
	}, nil

}

func (s *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	err := s.validator.Struct(req)
	if err != nil {
		return nil, err
	}

	res, err := s.userClient.FindUserByEmail(ctx, &userpb.FindUserByEmailRequest{
		Email: req.Email,
	})

	if err != nil {
		return nil, err
	}

	accToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   res.User.Id,
		"name": res.User.Username,
		"exp":  time.Now().Add(10 * time.Second).Unix(),
	})

	refToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   res.User.Id,
		"name": res.User.Username,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	})

	signedAccToken, err := accToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, err
	}
	signedRefToken, err := refToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	header := map[string]string{
		"access_token":  signedAccToken,
		"refresh_token": signedRefToken,
	}

	md := metadata.New(header)

	ctx = metadata.NewOutgoingContext(ctx, md)

	err = grpc.SetHeader(ctx, md)
	if err != nil {
		return nil, err
	}

	return &authpb.LoginResponse{
		Message: "success login",
	}, nil
}

func (s *AuthServer) Verify(ctx context.Context, _ *emptypb.Empty) (*authpb.VerifyResponse, error) {
	return &authpb.VerifyResponse{
		Message: "success verify",
	}, nil
}
