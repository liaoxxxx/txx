package controllers

import (
	"beego/orm"
	_ "crypto/md5"
	"fmt"
	"untitle_go_project/function"
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
	userName := c.GetString("user_name", "false")
	password := c.GetString("password", "false")

	if password == "false" {
		c.JsonFail("获取密码失败")
	}
	if userName == "false" {
		c.JsonFail("获取用户名失败")
	}
	o := orm.NewOrm()
	u := models.User{}

	u.User_name = userName

	err := o.Read(&u, "User_name")
	//如果未找到当前用户
	if err == orm.ErrNoRows {
		c.JsonFail("用户名出错，该用户名未注册")
		c.StopRun()
	}

	encPassword := function.GetEncPasswd(u.Salt, password)
	//加密后相等
	if encPassword == u.Password {
		c.SetSession("user_info", u)
		c.JsonSuccess("登录成功")
		c.StopRun()
	}

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
	u.Email = email

	o := orm.NewOrm()
	//用户名是否已注册
	err := o.Read(&u, "User_name")
	if err != orm.ErrNoRows {
		c.JsonFail("该用户名已注册，请更换")
	}
	//手机号码是否已注册
	err = o.Read(&u, "Phone")
	if err != orm.ErrNoRows {
		c.JsonFail("该手机号码已注册，请用本手机号码直接登录")
	}
	//邮箱是否已注册
	err = o.Read(&u, "Phone")
	if err != orm.ErrNoRows {
		c.JsonFail("该手机号码已邮箱，请用此邮箱直接登录")
	}
	//盐
	u.Salt = function.GetRandomString(5)
	//密码加密后
	u.Password = function.GetEncPasswd(u.Salt, password)

	id, err := o.Insert(&u)
	if err == nil {
		c.JsonSuccess(u.User_name + "您已注册成功为" + string(id) + "为用户")
		return
	}
	c.JsonFail("注册失败")
}
