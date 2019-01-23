package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"untitle_go_project/models"
	_ "untitle_go_project/routers"
)

func init() {
	dbType := beego.AppConfig.String("dbType")
	dbHost := beego.AppConfig.String("dbHost")
	dbName := beego.AppConfig.String("dbName")
	dbUrl := beego.AppConfig.String("dbUrl")
	dbPasaswd := beego.AppConfig.String("dbPasaswd")
	dbUser := beego.AppConfig.String("dbUser")
	connect := dbUser + ":" + dbPasaswd + "@tcp(" + dbUrl + ":" + dbHost + ")/" + dbName + "?charset=utf8"

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", dbType, connect, 30)
	orm.RegisterModel(new(models.User))
}

func main() {
	beego.Run()
}
