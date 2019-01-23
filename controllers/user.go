package controllers

import "fmt"

type UserController struct {
	indexBaseController
}

func (c *UserController) Prepare() {
}
func (c *UserController) Get() {
	c.TplName = "index/register/navbar.html"
}

func (c *UserController) Login() {
	c.TplName = "index/user/login.html"
}

func (c *UserController) Post() {
	userName := c.GetString("userName")
	fmt.Println(userName)
}
func (c *UserController) SignIn() {
	userName := c.GetString("userName")
	c.Data["userName"] = userName
	c.TplName = "index/user/sign.html"
}
