package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "test.html"
}

func (c *MainController) Post() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "test.html"
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["title"] = "管理员登录"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "admin/login/index.html"
}

func (c *LoginController) Post() {
	user_name := c.GetString("user_name")
	fmt.Println(user_name)
}




