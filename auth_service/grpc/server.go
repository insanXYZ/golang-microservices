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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

	signedAccToken, err := s.createJwtToken(jwt.MapClaims{
		"id":   res.User.Id,
		"name": res.User.Username,
		"exp":  time.Now().Add(10 * time.Minute).Unix(),
	})

	if err != nil {
		return nil, err
	}

	signedRefToken, err := s.createJwtToken(jwt.MapClaims{
		"id":   res.User.Id,
		"name": res.User.Username,
		"exp":  time.Now().Add(24 * time.Minute).Unix(),
	})

	if err != nil {
		return nil, err
	}

	header := map[string]string{
		"access_token":  signedAccToken,
		"refresh_token": signedRefToken,
	}

	md := metadata.New(header)

	err = grpc.SetHeader(ctx, md)
	if err != nil {
		return nil, err
	}

	return &authpb.LoginResponse{
		AccessToken:  signedAccToken,
		RefreshToken: signedRefToken,
	}, nil
}

func (s *AuthServer) Verify(ctx context.Context, _ *emptypb.Empty) (*authpb.VerifyResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}

	accToken := md.Get("access_token")[0]

	_, err := jwt.Parse(accToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Errorf(codes.PermissionDenied, "Unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	return &authpb.VerifyResponse{
		Message: "success verify",
	}, nil
}

func (s *AuthServer) Refresh(ctx context.Context, _ *emptypb.Empty) (*authpb.RefreshResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}

	refToken := md.Get("refresh_token")[0]

	token, err := jwt.Parse(refToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Errorf(codes.PermissionDenied, "Unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, "error assertion to jwt.mapclaims")
	}

	newAccToken, err := s.createJwtToken(jwt.MapClaims{
		"id":   claims["id"],
		"name": claims["name"],
		"exp":  time.Now().Add(10 * time.Minute).Unix(),
	})

	if err != nil {
		return nil, err
	}

	return &authpb.RefreshResponse{
		AccessToken: newAccToken,
	}, nil
}

func (s *AuthServer) createJwtToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}
