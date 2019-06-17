package handler

import (
	"rpc_blog/proto/user"
	"rpc_blog/src/user_srv/db"
)
import "context"

type NewUserServiceExtHandler struct {
}

func (u *NewUserServiceExtHandler) RegisterUser(ctx context.Context, req *user_ext.RegisterRequest, resp *user_ext.RegisterResponse) error {
	email := req.Email
	password := req.Password
	username := req.Username
	err := db.InsertUser(username, email, password)
	return err
}

func (u *NewUserServiceExtHandler) LoginUser(ctx context.Context, req *user_ext.RegisterRequest, resp *user_ext.RegisterResponse) error {
	email := req.Email
	password := req.Password
	user, err := db.GetUserByEmail(email, password)
	if err != nil {
		return err
	}

}
