package routers

import (
	"github.com/astaxie/beego"
	"untitle_go_project/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin/register", &controllers.RegisterController{})
	beego.Router("/user/register", &controllers.RegisterController{})
	beego.Router("/user/login", &controllers.UserController{}, "*:Login")
	beego.Router("/user/logout", &controllers.UserController{})

}
