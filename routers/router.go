package routers

import (
	"github.com/astaxie/beego"
	"untitle_go_project/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/registermanager", &controllers.RegisterController{})
}
