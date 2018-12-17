package routers

import (
	"txx/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/admin/login", &controllers.LoginController{})
}
