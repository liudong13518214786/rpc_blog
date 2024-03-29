package main

import (
	"github.com/astaxie/beego"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"rpc_blog/config"
	blog_ext "rpc_blog/proto/blog"
	"rpc_blog/src/blog_srv/db"
	"rpc_blog/src/blog_srv/handler"
	"time"
)

func main() {
	service := micro.NewService(micro.Name(config.NameSpace+config.ServiceNameBlog),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(20*time.Second),
		micro.Version("latest"),
	)

	service.Init(
		micro.Action(func(context *cli.Context) {
			db.InitDatabase(config.PostgresqlDSN_Blog)
			_ = blog_ext.RegisterBlogServiceHandler(service.Server(),
				new(handler.BlogServiceExtHandler),
				server.InternalHandler(true))
		}),
		micro.AfterStop(func() error {
			beego.Info("博客服务已经停止")
			return nil
		}),
		micro.AfterStart(func() error {
			beego.Info("博客服务已启动...")
			return nil
		}),
	)
	if err := service.Run(); err != nil {
		//
		beego.Error("博客服务启动失败:", err)
	}
}
