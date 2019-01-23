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

func (c *UserController) SignIn() {
	username := c.GetString("username")
	password := c.GetString("password")

	c.Data["username"] = username
	c.Data["password"] = password
}

func (c *UserController) Post() {
	userName := c.GetString("userName")
	fmt.Println(userName)
}

//register user

func (c *UserController) Add() {
	c.TplName = "index/user/add.html"
}

func (c *UserController) Register() {
	//todo method not allowed
}
