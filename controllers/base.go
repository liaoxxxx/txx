package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"io"
)

type baseController struct {
	beego.Controller
	orm           orm.Ormer
	moduleName    string
	contrllerName string
	actionName    string
}

//beego.controller [T] 的 init method
func (c *baseController) Prepare() {
	c.moduleName = "admin"                                     //定义模块名为 admin
	c.contrllerName, c.actionName = c.GetControllerAndAction() //获取abc 的 contrllerName,actionName赋值到adminBaseController 的 contrllerName |	actionName
	c.orm = orm.NewOrm()                                       //
}
func (c *baseController) JsonSuccess(msg string, data ...interface{}) {
	jsonData := make(map[string]interface{}, 4)
	jsonData["status"] = true
	jsonData["msg"] = msg
	jsonData["code"] = 200
	jsonData["data"] = data
	returnJSON, err := json.Marshal(jsonData)
	if err != nil {
		beego.Error(err)
	}
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Ctx.ResponseWriter.Header().Set("Cache-Control", "no-cache, no-store")
	io.WriteString(c.Ctx.ResponseWriter, string(returnJSON))
	c.StopRun()
}

func (c *baseController) JsonFail(msg string, data ...interface{}) {
	jsonData := make(map[string]interface{}, 4)
	jsonData["status"] = true
	jsonData["msg"] = msg
	jsonData["code"] = 200
	jsonData["data"] = data
	returnJSON, err := json.Marshal(jsonData)
	if err != nil {
		beego.Error(err)
	}
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Ctx.ResponseWriter.Header().Set("Cache-Control", "no-cache, no-store")
	io.WriteString(c.Ctx.ResponseWriter, string(returnJSON))
	c.StopRun()
}
