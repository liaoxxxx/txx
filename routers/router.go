package routers

import (
	"github.com/astaxie/beego"
	"untitle_go_project/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin/login", &controllers.LoginController{})
	beego.Router("/addmanager", &controllers.ManagerController{})
	beego.Router("/addmanager", &controllers.ManagerController{})
	beego.Router("/registermanager", &controllers.RegisterController{})
}
