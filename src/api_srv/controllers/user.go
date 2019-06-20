package controllers

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/micro/go-micro"
	"rpc_blog/config"
	user_ext "rpc_blog/proto/user"
	"rpc_blog/src/module/utils"
)

type UserControllers struct {
	beego.Controller
}

func (u *UserControllers) Post() {
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
	service := micro.NewService(micro.Name(config.NameSpace + config.ServiceNameUser + "client"))
	service.Init()
	userService := user_ext.NewUserService(config.NameSpace+config.ServiceNameUser, service.Client())
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
