package routers

import (
	"github.com/astaxie/beego"
	"rpc_blog/src/api_srv/controllers"
)

func init() {
	beego.Router("/", &controllers.BlogControllers{})
	beego.Router("/register", &controllers.UserRegisterControllers{})
	beego.Router("/login", &controllers.UserControllers{})
}
