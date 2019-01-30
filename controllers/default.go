package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	err := c.GetSession("user_info")
	if err != nil {
		fmt.Println(err)
	}
	c.TplName = "index/index/index.html"
}

func (c *MainController) Post() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
}
