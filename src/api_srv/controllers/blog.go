package controllers

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/micro/go-micro"
	"rpc_blog/config"
	blog_ext "rpc_blog/proto/blog"
)

type BlogControllers struct {
	beego.Controller
}

func (b *BlogControllers) Get() {
	service := micro.NewService(micro.Name(config.NameSpace + config.ServiceNameBlog + "client"))
	service.Init()
	blogserver := blog_ext.NewBlogService(config.NameSpace+config.ServiceNameBlog, service.Client())
	blogResponse, err := blogserver.GetBlogList(context.TODO(), &blog_ext.BlogListRequest{Userid: "123"})
	if err != nil {
		beego.Error("博客服务：获取博客列表失败：", err)
		return
	}
	a := blogResponse.BlogList
	beego.Info(a)
	b.Ctx.WriteString("123123")

}
