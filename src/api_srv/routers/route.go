package routers

import (
	"github.com/astaxie/beego"
	"rpc_blog/src/api_srv/controllers"
)

func init() {
	//登录注册相关
	beego.Router("/register", &controllers.UserRegisterControllers{})
	beego.Router("/login", &controllers.UserLoginControllers{})
	beego.Router("/logout", &controllers.UserLogOutControllers{})

	//博客相关
	beego.Router("/", &controllers.BlogControllers{})
	beego.Router("/write", &controllers.BlogWriteControllers{})
}
