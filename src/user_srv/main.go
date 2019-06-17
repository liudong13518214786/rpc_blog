package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"rpc_blog/config"
	user_2 "rpc_blog/proto/user"
	"rpc_blog/src/user_srv/db"
)

func main() {
	service := micro.NewService(micro.Name(config.NameSpace+config.ServiceNameUser),
		micro.Version("latest"))
	//定义service的操作
	service.Init(
		micro.Action(func(context *cli.Context) {
			db.InitDatabase(config.MysqlDSN)
			_ = user_2.RegisterUserServiceHandler(service.Server(), user_2.UserServiceHandler(), server.InternalHandler(true))
		}),
	)
}
