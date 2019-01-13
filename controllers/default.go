package controllers

import (
	"beego/orm"
	"github.com/astaxie/beego"
	"untitle_go_project1/models/user"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//new 一个orm 对象
	O := orm.NewOrm()
	//new model
	u := user.User{}
	//赋值
	u.Nick_name = "廖13"
	u.User_name = "廖永坚"
	u.Passwd = "liao993501756"
	u.Salt = "185454"
	//插入
	_, err := O.Insert(&u)
	if err != nil {
		beego.Info("插入失败", &err)
		return
	}
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

}

type ManagerController struct {
	beego.Controller
}

func (C *ManagerController) Get() {
	//orm 对象
	o := orm.NewOrm()
	o.Using("default")
	//model 结构体对象
	u := user.User{}
	//赋值
	u.Name = "liaoxx124"
	//查询 如果根据id 外的字段查询，传入 param1 为 相应的model结构体 对象 ，不是这个对象的 指针类型 ，param2 为字段名 首字母大写
	err := o.Read(u, "Name") //将会把查询结果赋值回 该model结构体对象
	if err != nil {
		beego.Info("查询失败", err)
	} else {
		beego.Info("查询成功", u.Passwd)
	}
	return
}
