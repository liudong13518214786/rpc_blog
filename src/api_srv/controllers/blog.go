package controllers

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/micro/go-micro"
	"rpc_blog/config"
	blog_ext "rpc_blog/proto/blog"
	user_ext "rpc_blog/proto/user"
	"rpc_blog/src/module/utils"
)

type BlogControllers struct {
	beego.Controller
}

type BlogWriteControllers struct {
	beego.Controller
}

type BlogDeleteControllers struct {
	beego.Controller
}

type BlogDetailControllers struct {
	beego.Controller
}

type BlogModifyControllers struct {
	beego.Controller
}

func initBlogService() blog_ext.BlogService {
	service := micro.NewService(micro.Name(config.NameSpace + config.ServiceNameBlog + "client"))
	service.Init()
	blogserver := blog_ext.NewBlogService(config.NameSpace+config.ServiceNameBlog, service.Client())
	return blogserver
}

func (b *BlogControllers) Get() {
	//service := micro.NewService(micro.Name(config.NameSpace + config.ServiceNameBlog + "client"))
	//service.Init()
	//blogserver := blog_ext.NewBlogService(config.NameSpace+config.ServiceNameBlog, service.Client())
	blogserver := initBlogService()
	blogResponse, err := blogserver.GetBlogList(context.TODO(), &blog_ext.BlogListRequest{Userid: "123"})
	if err != nil {
		beego.Error("博客服务：获取博客列表失败：", err)
		return
	}
	bloginfo := blogResponse.BlogList
	b.Data["json"] = map[string]interface{}{"code": 100, "message": bloginfo}
	b.ServeJSON()
	return

}

func (b *BlogWriteControllers) Post() {
	bid := b.GetString("bid")
	token := b.GetString("token")
	title := b.GetString("title")
	info := b.GetString("info")
	result, success := utils.CheckParma(token, title, info)
	if !success {
		b.Data["json"] = map[string]interface{}{"code": 500, "message": result}
		b.ServeJSON()
		return
	}
	userService := initUserService()
	verifyResponse, err := userService.VerifyToken(context.TODO(), &user_ext.VerifyRequest{Token: token})
	if err != nil {
		beego.Error("用户服务：校验用户token失败", err)
		b.Data["json"] = map[string]interface{}{"code": 500, "message": "校验用户token失败"}
		b.ServeJSON()
		return
	}
	if verifyResponse.Code != 100 {
		b.Data["json"] = map[string]interface{}{"code": verifyResponse.Code, "message": verifyResponse.Message}
		b.ServeJSON()
		return
	}
	userId := verifyResponse.Message
	blogService := initBlogService()
	writeResponse, err := blogService.WriteBlog(context.TODO(), &blog_ext.WriteBlogRequest{BlogId: bid, Userid: userId, Title: title, Info: info})
	if err != nil {
		beego.Error("博客服务：写博客失败")
		b.Data["json"] = map[string]interface{}{"code": 500, "message": "write error"}
		b.ServeJSON()
		return
	}
	b.Data["json"] = map[string]interface{}{"code": writeResponse.Code, "message": writeResponse.Message}
	b.ServeJSON()
	return
}

func (b *BlogDeleteControllers) Post() {
	bid := b.GetString("bid")
	token := b.GetString("token")
	result, status := utils.CheckParma(bid, token)
	if !status {
		b.Data["json"] = map[string]interface{}{"code": 500, "message": result}
		b.ServeJSON()
		return
	}
	userService := initUserService()
	verifyResponse, err := userService.VerifyToken(context.TODO(), &user_ext.VerifyRequest{Token: token})
	if err != nil {
		beego.Error("用户服务：验证token失败")
		b.Data["json"] = map[string]interface{}{"code": 500, "message": "token 错误"}
		b.ServeJSON()
		return
	}
	if verifyResponse.Code != 100 {
		b.Data["json"] = map[string]interface{}{"code": verifyResponse.Code, "message": verifyResponse.Message}
		b.ServeJSON()
		return
	}
	userId := verifyResponse.Message
	blogService := initBlogService()
	deleteResponse, err := blogService.DeleteBlog(context.TODO(), &blog_ext.DeleteBlogRequest{BlogId: bid, Userid: userId})
	if err != nil {
		beego.Error("博客服务：删除博客失败")
		b.Data["json"] = map[string]interface{}{"code": 500, "message": "删除博客失败"}
		b.ServeJSON()
		return
	}
	b.Data["json"] = map[string]interface{}{"code": deleteResponse.Code, "message": deleteResponse.Message}
	b.ServeJSON()
	return
}

func (b *BlogDetailControllers) Get() {
	bid := b.GetString("bid")
	result, status := utils.CheckParma(bid)
	if !status {
		b.Data["json"] = map[string]interface{}{"code": 500, "message": result}
		b.ServeJSON()
		return
	}
	BlogService := initBlogService()
	blogDetailResponse, err := BlogService.BlogDetail(context.TODO(), &blog_ext.BlogDetailRequest{BlogId: bid})
	if err != nil {
		beego.Error("博客服务：请求博客详情失败")
		b.Data["json"] = map[string]interface{}{"code": 500, "message": "请求博客详情失败"}
		b.ServeJSON()
		return
	}
	b.Data["json"] = map[string]interface{}{"code": 500, "message": blogDetailResponse}
	b.ServeJSON()
	return

}

func (b *BlogModifyControllers) Post() {
	bid := b.GetString("bid")
	title := b.GetString("title")
	detail := b.GetString("detail")
	Token := b.GetString("token")
	result, status := utils.CheckParma(bid, title, detail, Token)
	if !status {
		b.Data["json"] = map[string]interface{}{"code": 500, "message": result}
		b.ServeJSON()
		return
	}
	UserService := initUserService()
	verifyTokenResponse, err := UserService.VerifyToken(context.TODO(), &user_ext.VerifyRequest{Token: Token})
	if err != nil {
		beego.Error("用户服务：验证码token失败")
		b.Data["json"] = map[string]interface{}{"code": 500, "message": "验证码token失败"}
		b.ServeJSON()
		return
	}
	if verifyTokenResponse.Code != 100 {
		b.Data["json"] = map[string]interface{}{"code": verifyTokenResponse.Code, "message": verifyTokenResponse.Message}
		b.ServeJSON()
		return
	}
	userId := verifyTokenResponse.Message
	BlogService := initBlogService()
	ModifyBlogResponse, err := BlogService.ModifyBlog(context.TODO(), &blog_ext.ModifyBlogRequest{BlogId: bid, Title: title, Detail: detail, Userid: userId})
	if err != nil {
		beego.Error("博客服务：修改博客失败")
		b.Data["json"] = map[string]interface{}{"code": 500, "message": "修改失败"}
		b.ServeJSON()
		return
	}
	b.Data["json"] = map[string]interface{}{"code": ModifyBlogResponse.Code, "message": ModifyBlogResponse.Message}
	b.ServeJSON()
	return
}
