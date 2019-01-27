package controllers

import (
	"beego/orm"
	"fmt"
	"untitle_go_project/models"
)

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
	userName := c.GetString("userName", "false")
	password := c.GetString("password", "false")
	email := c.GetString("email", "false")
	phone := c.GetString("phone", "false")
	if userName == "false" {
		c.JsonFail("获取用户名失败")
	}
	if password == "false" {
		c.JsonFail("获取密码失败")
	}
	if email == "false" {
		c.JsonFail("获取邮箱失败")
	}
	if phone == "false" {
		c.JsonFail("获取用户手机号码失败")
	}
	u := models.User{}
	u.User_name = userName
	u.Phone = phone

	o := orm.NewOrm()
	err := o.Read(&u, "Phone")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
	c.JsonSuccess(string(u.Full_name))
}
