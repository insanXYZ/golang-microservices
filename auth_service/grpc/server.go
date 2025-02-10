package main

import (
	"context"

	"github.com/go-playground/validator/v10"
	authpb "github.com/insanXYZ/proto/gen/go/auth"
	userpb "github.com/insanXYZ/proto/gen/go/user"
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

func (s *AuthServer) Login(context.Context, *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	return nil, nil
}
