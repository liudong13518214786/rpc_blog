package handler

import (
	"errors"
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
		err := errors.New(res)
		return err
	}
	user, _ := db.GetUserByEmail(email)
	if user != nil {
		resp.Code = 512
		resp.Message = "Email has been registered"
		err := errors.New("Email has been registered")
		return err
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
	user, err := db.GetUserByEmailPassword(email, password)
	if err != nil {
		resp.Code = 510
		resp.Message = "error email"
		return err
	}
	if user == nil {
		resp.Code = 511
		resp.Message = "error email"
		return err
	}
	sdtoken := utils.CreateToken(user.UserId)
	cache.Mcache.Delete(sdtoken)
	resx := cache.Mcache.Set("token", sdtoken, 86400)
	if resx {
		resp.Code = 100
		resp.Message = sdtoken
		return nil
	}
	resp.Code = 506
	resp.Message = "登录失败"
	return nil
}

func (u *NewUserServiceExtHandler) LogOutUser(ctx context.Context, req *user_ext.LogOutRequest, resp *user_ext.LogOutResponse) error {
	token := req.Token
	res, success := utils.CheckParma(token)
	if !success {
		resp.Code = 505
		resp.Message = res
		e := errors.New("token err")
		return e
	}
	_, err := utils.CheckToken(token)
	if err != nil {
		resp.Code = 500
		resp.Message = "token error"
		return err
	}
	cache.Mcache.Delete(token)
	resp.Code = 100
	resp.Message = "logout success"
	return nil
}

func (u *NewUserServiceExtHandler) VerifyToken(ctx context.Context, req *user_ext.VerifyRequest, resp *user_ext.VerifyResponse) error {
	token := req.Token
	res, err := utils.CheckToken(token)
	if err != nil {
		resp.Code = 500
		resp.Message = res
		return err
	}
	userid := res
	user, err := db.GetUserByUuid(userid)
	if err != nil {
		resp.Code = 505
		resp.Message = "token error"
		return err
	}
	resp.Code = 100
	resp.Message = user.UserId
	return nil
}
