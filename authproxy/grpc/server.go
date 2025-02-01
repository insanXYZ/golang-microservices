package main

import (
	"context"

	"github.com/go-playground/validator/v10"
	authsv "github.com/insanXYZ/proto/gen/go/auth"
	usersv "github.com/insanXYZ/proto/gen/go/user"
)

type AuthServer struct {
	userClient usersv.UserServiceClient
	validator  *validator.Validate
	authsv.UnimplementedAuthServiceServer
}

func NewAuthServer(userClient usersv.UserServiceClient, validator *validator.Validate) *AuthServer {
	return &AuthServer{
		userClient: userClient,
		validator:  validator,
	}
}

func (s *AuthServer) Register(ctx context.Context, req *authsv.RegisterRequest) (*authsv.RegisterResponse, error) {
	err := s.validator.Struct(req)
	if err != nil {
		return nil, err
	}

	res, err := s.userClient.Insert(ctx, &usersv.InsertRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})

	if err != nil {
		return nil, err
	}

	return &authsv.RegisterResponse{
		Message: res.Message,
	}, nil

}

func (s *AuthServer) Login(context.Context, *authsv.LoginRequest) (*authsv.LoginResponse, error) {
	return nil, nil
}
