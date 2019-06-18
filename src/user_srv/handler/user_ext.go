package handler

import (
	"rpc_blog/proto/user"
	"rpc_blog/src/module/cache"
	"rpc_blog/src/module/utils"
	"rpc_blog/src/user_srv/db"
)
import "context"

type NewUserServiceExtHandler struct {
}

//func NewUserServiceHandler() *NewUserServiceExtHandler {
//	return &NewUserServiceExtHandler{}
//}

func (u *NewUserServiceExtHandler) RegisterUser(ctx context.Context, req *user_ext.RegisterRequest, resp *user_ext.RegisterResponse) error {
	email := req.Email
	password := req.Password
	username := req.Username
	res, success := utils.CheckParma(email, password, username)
	if !success {
		resp.Code = 505
		resp.Message = res
		return nil
	}
	err := db.InsertUser(username, email, password)
	if err != nil {
		resp.Code = 500
		resp.Message = "user register failed"
	} else {
		resp.Code = 100
		resp.Message = "user register success"
	}
	return err
}

func (u *NewUserServiceExtHandler) LoginUser(ctx context.Context, req *user_ext.LoginRequest, resp *user_ext.LoginResponse) error {
	email := req.Email
	password := req.Password
	res, success := utils.CheckParma(email, password)
	if !success {
		resp.Code = 505
		resp.Message = res
		return nil
	}
	user, err := db.GetUserByEmail(email, password)
	if err != nil {
		return err
	}
	if user == nil {
		return err
	}
	sdtoken := utils.CreateToken(email)
	cache.Mcache.Delete(sdtoken)
	resx := cache.Mcache.Set("token", sdtoken, 86400)
	if resx {
		resp.Code = 100
		resp.Message = "login success"
	}
	return nil
}

func (u *NewUserServiceExtHandler) LogOutUser(ctx context.Context, req *user_ext.LogOutRequest, resp *user_ext.LogOutResponse) error {
	token := req.Token
	res, success := utils.CheckParma(token)
	if !success {
		resp.Code = 505
		resp.Message = res
		return nil
	}
	_, err := utils.CheckToken(token)
	if err != nil {
		resp.Code = 500
		resp.Message = "token error"
		return nil
	}
	cache.Mcache.Delete(token)
	return nil
}
