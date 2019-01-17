package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	fmt.Println("测试")
}
