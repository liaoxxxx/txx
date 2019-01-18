package controllers

import (
	"fmt"
)

type RegisterController struct {
	adminBaseController
}

func (c *RegisterController) Prepare() {
	c.adminBaseController.Prepare()
}
func (c *RegisterController) Get() {
	fmt.Println("测试")
}
