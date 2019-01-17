package controllers

import (
	"beego"
	"beego/orm"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type adminBaseController struct {
	beego.Controller
	orm           orm.Ormer
	moduleName    string
	contrllerName string
	actionName    string
}

//beego.controller [T] 的 init method
func (abc *adminBaseController) Prepare() {
	abc.moduleName = "admin"                                         //定义模块名为 admin
	abc.contrllerName, abc.actionName = abc.GetControllerAndAction() //获取abc 的 contrllerName,actionName赋值到adminBaseController 的 contrllerName |	actionName
	abc.orm = orm.NewOrm()                                           //
}
