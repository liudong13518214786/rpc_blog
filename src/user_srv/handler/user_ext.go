package handler

import (
	"rpc_blog/proto/user"
	"rpc_blog/src/user_srv/db"
)
import "context"

type NewUserServiceExtHandler struct {
}

func NewUserServiceHandler() *NewUserServiceExtHandler {
	return &NewUserServiceExtHandler{}
}

func (u *NewUserServiceExtHandler) RegisterUser(ctx context.Context, req *user_ext.RegisterRequest, resp *user_ext.RegisterResponse) error {
	email := req.Email
	password := req.Password
	username := req.Username
	err := db.InsertUser(username, email, password)
	return err
}

func (u *NewUserServiceExtHandler) LoginUser(ctx context.Context, req *user_ext.LoginRequest, resp *user_ext.LoginResponse) error {
	email := req.Email
	password := req.Password
	user, err := db.GetUserByEmail(email, password)
	if err != nil {
		return err
	}
	if user == nil {
		return err
	}
	//todo 登录处理
	return nil
}

func (u *NewUserServiceExtHandler) LogOutUser(ctx context.Context, req *user_ext.LogOutRequest, resp *user_ext.LogOutResponse) error {
	return nil
}
