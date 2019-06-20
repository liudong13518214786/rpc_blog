package controllers

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/micro/go-micro"
	"rpc_blog/config"
	user_ext "rpc_blog/proto/user"
	"rpc_blog/src/module/utils"
)

type UserRegisterControllers struct {
	beego.Controller
}

type UserLoginControllers struct {
	beego.Controller
}

func initUserService() user_ext.UserService {
	service := micro.NewService(micro.Name(config.NameSpace + config.ServiceNameUser + "client"))
	service.Init()
	user_service := user_ext.NewUserService(config.NameSpace+config.ServiceNameUser, service.Client())
	return user_service
}

func (u *UserRegisterControllers) Post() {
	//注册用户
	email := u.GetString("email")
	password := u.GetString("password")
	usename := u.GetString("username")
	res, status := utils.CheckParma(email, password, usename)
	if !status {
		u.Data["json"] = map[string]interface{}{"code": 500, "message": res}
		u.ServeJSON()
		return
	}
	//service := micro.NewService(micro.Name(config.NameSpace + config.ServiceNameUser + "client"))
	//service.Init()
	//userService := user_ext.NewUserService(config.NameSpace+config.ServiceNameUser, service.Client())
	userService := initUserService()
	registerResponse, err := userService.RegisterUser(context.TODO(), &user_ext.RegisterRequest{Email: email, Username: usename, Password: password})
	if err != nil {
		beego.Error("用户服务：注册用户失败")
		u.Data["json"] = map[string]interface{}{"code": 520, "message": err}
		u.ServeJSON()
		return
	}
	u.Data["json"] = map[string]interface{}{"code": registerResponse.Code, "message": registerResponse.Message}
	u.ServeJSON()
	return

}

func (u *UserLoginControllers) Post() {
	email := u.GetString("email")
	password := u.GetString("password")
	res, status := utils.CheckParma(email, password)
	if !status {
		u.Data["json"] = map[string]interface{}{"code": 500, "message": res}
		u.ServeJSON()
		return
	}
	userService := initUserService()
	response, err := userService.LoginUser(context.TODO(), &user_ext.LoginRequest{Email: email, Password: password})
	if err != nil {
		beego.Error("用户服务：登录用户失败", err)
		u.Data["json"] = map[string]interface{}{"code": 500, "message": "用户服务：登录用户失"}
		u.ServeJSON()
		return
	}
	u.Data["json"] = map[string]interface{}{"code": response.Code, "message": response.Message}
	u.ServeJSON()
	return
}
