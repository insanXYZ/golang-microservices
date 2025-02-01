package main

import (
	"context"
	"encoding/json"
	"fmt"

	usersv "github.com/insanXYZ/proto/gen/go/user"
)

type UserServer struct {
	usersv.UnimplementedUserServiceServer
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (u *UserServer) Insert(ctx context.Context, req *usersv.InsertRequest) (*usersv.InsertResponse, error) {
	b, err := json.MarshalIndent(req, "", " ")
	if err != nil {
		return nil, err
	}
	fmt.Println(string(b))

	return &usersv.InsertResponse{
		Message: "request insert success",
	}, nil
}

// func (u *UserServer) FindUserByEmail(context.Context, *usersv.FindUserByEmailRequest) (*usersv.FindUserByEmailResponse, error) {

// }
