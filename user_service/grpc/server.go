package main

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	authpb "github.com/insanXYZ/proto/gen/go/auth"
	userpb "github.com/insanXYZ/proto/gen/go/user"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserServer struct {
	db         *pgx.Conn
	validator  *validator.Validate
	authClient authpb.AuthServiceClient
	userpb.UnimplementedUserServiceServer
}

func NewUserServer(db *pgx.Conn, authClient authpb.AuthServiceClient, validator *validator.Validate) *UserServer {
	return &UserServer{
		db:         db,
		validator:  validator,
		authClient: authClient,
	}
}

func (u *UserServer) Insert(ctx context.Context, req *userpb.InsertRequest) (*userpb.InsertResponse, error) {
	LogPrintln("using insert rpc")
	err := u.validator.Struct(req)
	if err != nil {
		LogPrintln("Error validation struct", err.Error())
		return nil, err
	}

	tx, err := u.db.Begin(ctx)
	if err != nil {
		LogPrintln("Error transaction database", err.Error())
		return nil, err
	}

	defer tx.Rollback(ctx)

	passByte, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		LogPrintln("Error generate bcrypt", err.Error())
		return nil, err
	}

	_, err = tx.Exec(ctx, "insert into users(id, username, email, password) values($1,$2,$3,$4)", uuid.NewString(), req.Username, req.Email, string(passByte))
	if err != nil {
		LogPrintln("Error exec command", err.Error())
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		LogPrintln("Error commit transaction", err.Error())
		return nil, err
	}

	return &userpb.InsertResponse{
		Message: "success create user",
	}, nil

}

func (u *UserServer) FindUserByEmail(ctx context.Context, req *userpb.FindUserByEmailRequest) (*userpb.FindUserByEmailResponse, error) {
	LogPrintln("using finduserbyemail rpc")
	err := u.validator.Struct(req)
	if err != nil {
		LogPrintln("Error validation struct", err.Error())
		return nil, err
	}

	var user userpb.User

	err = u.db.QueryRow(ctx, "select id, username, email from users where email = $1", req.Email).Scan(&user.Id, &user.Username, &user.Email)

	if err != nil {
		LogPrintln("Error exec query", err.Error())
		return nil, err
	}

	return &userpb.FindUserByEmailResponse{
		User: &user,
	}, nil
}
