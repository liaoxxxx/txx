package routers

import (
	"untitle_go_project1/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/admin/login", &controllers.LoginController{})
    beego.Router("/addmanager",&controllers.ManagerController{})
}
