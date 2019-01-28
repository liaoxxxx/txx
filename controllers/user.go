package controllers

import (
	"beego/orm"
	"crypto/md5"
	_ "crypto/md5"
	"encoding/hex"
	"fmt"
	"untitle_go_project/func"
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
	userName := c.GetString("userName", "false")
	password := c.GetString("password", "false")
	email := c.GetString("email", "false")
	phone := c.GetString("phone", "false")

	h := md5.New()
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
	u.Salt = _func.GetRandomString(5)
	h.Write([]byte(u.Salt + password))
	//密码加密后
	u.Password = hex.EncodeToString(h.Sum(nil))

	id, err := o.Insert(&u)
	if err == nil {
		c.JsonSuccess(u.User_name + "您已注册成功为" + string(id) + "为用户")
		return
	}
	c.JsonFail("注册失败")
}
