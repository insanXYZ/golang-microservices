package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	authpb "github.com/insanXYZ/proto/gen/go/auth"
	chatpb "github.com/insanXYZ/proto/gen/go/chat"
	userpb "github.com/insanXYZ/proto/gen/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthServer struct {
	userClient userpb.UserServiceClient
	chatClient chatpb.ChatServiceClient
	validator  *validator.Validate
	authpb.UnimplementedAuthServiceServer
}

func NewAuthServer(userClient userpb.UserServiceClient, chatClient chatpb.ChatServiceClient, validator *validator.Validate) *AuthServer {
	return &AuthServer{
		userClient: userClient,
		chatClient: chatClient,
		validator:  validator,
	}
}

func (s *AuthServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	LogPrintln("using register rpc")
	err := s.validator.Struct(req)
	if err != nil {
		LogPrintln("Error validation struct", err.Error())
		return nil, err
	}

	res, err := s.userClient.Insert(ctx, &userpb.InsertRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})

	if err != nil {
		LogPrintln("Error Insert to User service", err.Error())
		return nil, err
	}

	return &authpb.RegisterResponse{
		Message: res.Message,
	}, nil

}

func (s *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	LogPrintln("using login rpc")
	err := s.validator.Struct(req)
	if err != nil {
		LogPrintln("Error validation struct", err.Error())
		return nil, err
	}

	res, err := s.userClient.FindUserByEmail(ctx, &userpb.FindUserByEmailRequest{
		Email: req.Email,
	})

	if err != nil {
		LogPrintln("Error find user by email to User service", err.Error())
		return nil, err
	}

	signedAccToken, err := s.createJwtToken(jwt.MapClaims{
		"id":   res.User.Id,
		"name": res.User.Username,
		"exp":  time.Now().Add(10 * time.Minute).Unix(),
	})

	if err != nil {
		LogPrintln("Error create access token", err.Error())
		return nil, err
	}

	signedRefToken, err := s.createJwtToken(jwt.MapClaims{
		"id":   res.User.Id,
		"name": res.User.Username,
		"exp":  time.Now().Add(24 * time.Minute).Unix(),
	})

	if err != nil {
		LogPrintln("Error create refresh token", err.Error())
		return nil, err
	}

	header := map[string]string{
		"access-token":  signedAccToken,
		"refresh-token": signedRefToken,
	}

	md := metadata.New(header)

	err = grpc.SetHeader(ctx, md)
	if err != nil {
		LogPrintln("Error setheader grpc", err.Error())
		return nil, err
	}

	return &authpb.LoginResponse{
		AccessToken:  signedAccToken,
		RefreshToken: signedRefToken,
	}, nil
}

func (s *AuthServer) Verify(ctx context.Context, _ *emptypb.Empty) (*authpb.VerifyResponse, error) {
	LogPrintln("using verify rpc")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		LogPrintln("Error get metadata with fromIncomingContext")
		return nil, status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}
	fmt.Println(md)
	accTokens := md.Get("access-token")
	if len(accTokens) == 0 {
		LogPrintln("access-token are missing")
		return nil, status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}

	token, err := jwt.Parse(accTokens[0], func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Errorf(codes.PermissionDenied, "Unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		log.Println("Error parse jwt", err.Error())
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)

	return &authpb.VerifyResponse{
		User: &userpb.User{
			Id:       claims["id"].(string),
			Username: claims["username"].(string),
			Email:    claims["email"].(string),
		},
	}, nil
}

func (s *AuthServer) Refresh(ctx context.Context, _ *emptypb.Empty) (*authpb.RefreshResponse, error) {
	LogPrintln("using refresh rpc")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		LogPrintln("Error get metadata with fromIncomingContext")
		return nil, status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}

	refTokens := md.Get("refresh-token")

	if len(refTokens) == 0 {
		LogPrintln("access-token are missing")
		return nil, status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}

	token, err := jwt.Parse(refTokens[0], func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Errorf(codes.PermissionDenied, "Unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		log.Println("Error parse jwt", err.Error())
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		LogPrintln("Error assertion claims jwt")
		return nil, status.Error(codes.PermissionDenied, "error assertion to jwt.mapclaims")
	}

	newAccToken, err := s.createJwtToken(jwt.MapClaims{
		"id":   claims["id"],
		"name": claims["name"],
		"exp":  time.Now().Add(10 * time.Minute).Unix(),
	})

	if err != nil {
		LogPrintln("Error create new access token jwt", err.Error())
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
