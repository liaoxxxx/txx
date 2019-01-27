package controllers

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type adminBaseController struct {
	baseController
	orm           orm.Ormer
	moduleName    string
	contrllerName string
	actionName    string
}

//beego.controller [T] 的 init method
func (c *adminBaseController) Prepare() {
	c.moduleName = "admin"                                     //定义模块名为 admin
	c.contrllerName, c.actionName = c.GetControllerAndAction() //获取abc 的 contrllerName,actionName赋值到adminBaseController 的 contrllerName |	actionName
	c.orm = orm.NewOrm()                                       //
}
