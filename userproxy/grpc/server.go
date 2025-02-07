package main

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	usersv "github.com/insanXYZ/proto/gen/go/user"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserServer struct {
	db        *pgx.Conn
	validator *validator.Validate
	usersv.UnimplementedUserServiceServer
}

func NewUserServer(db *pgx.Conn, validator *validator.Validate) *UserServer {
	return &UserServer{
		db:        db,
		validator: validator,
	}
}

func (u *UserServer) Insert(ctx context.Context, req *usersv.InsertRequest) (*usersv.InsertResponse, error) {
	err := u.validator.Struct(req)
	if err != nil {
		return nil, err
	}

	tx, err := u.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	passByte, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(ctx, "insert into users(id, username, email, password) values($1,$2,$3,$4)", uuid.NewString(), req.Username, req.Email, string(passByte))
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &usersv.InsertResponse{
		Message: "success create user",
	}, nil

}

func (u *UserServer) FindUserByEmail(ctx context.Context, req *usersv.FindUserByEmailRequest) (*usersv.FindUserByEmailResponse, error) {
	err := u.validator.Struct(req)
	if err != nil {
		return nil, err
	}

	var user usersv.UserModel

	err = u.db.QueryRow(ctx, "select id, username, email, password from users where email = $1", req.Email).Scan(&user.Id, &user.Username, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &usersv.FindUserByEmailResponse{
		User: &user,
	}, nil
}
