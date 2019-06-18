package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"rpc_blog/config"
	user_ext "rpc_blog/proto/user"
	"rpc_blog/src/user_srv/db"
	"rpc_blog/src/user_srv/handler"
)

func main() {
	service := micro.NewService(micro.Name(config.NameSpace+config.ServiceNameUser),
		micro.Version("latest"))
	//定义service的操作
	service.Init(
		micro.Action(func(context *cli.Context) {
			db.InitDatabase(config.MysqlDSN)
			_ = user_ext.RegisterUserServiceHandler(service.Server(), new(handler.NewUserServiceExtHandler), server.InternalHandler(true))
		}),
		micro.AfterStop(func() error {
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)
	if err := service.Run(); err != nil {
		//启动失败

	}
}
