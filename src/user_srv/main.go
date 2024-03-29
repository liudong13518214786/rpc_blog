package main

import (
	"github.com/astaxie/beego"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"rpc_blog/config"
	user_ext "rpc_blog/proto/user"
	"rpc_blog/src/user_srv/db"
	"rpc_blog/src/user_srv/handler"
	"time"
)

func main() {
	service := micro.NewService(micro.Name(config.NameSpace+config.ServiceNameUser),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),      //注册服务的过期时间
		micro.RegisterInterval(time.Second*20)) //间隔多久再次注册服务
	//micro.Registry(server.Registry())
	//定义service的操作
	service.Init(
		micro.Action(func(context *cli.Context) {
			db.InitDatabase(config.PostgresqlDSN_User)
			_ = user_ext.RegisterUserServiceHandler(service.Server(), new(handler.NewUserServiceExtHandler),
				server.InternalHandler(true))
		}),
		micro.AfterStop(func() error {
			beego.Info("用户服务已经停止")
			return nil
		}),
		micro.AfterStart(func() error {
			beego.Info("用户服务正在运行...")
			return nil
		}),
	)

	if err := service.Run(); err != nil {
		//启动失败
		beego.Error("用户服务启动失败")
	}
}
