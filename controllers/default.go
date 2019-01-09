package controllers

import (
	"beego/orm"
	"fmt"
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

}

func insert() {

	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	user := user.User{User_name: "wood"}

	fmt.Println(o.Insert(&user))
}
