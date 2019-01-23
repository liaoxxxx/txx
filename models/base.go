package models

import (
	"beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {

	dbType := beego.AppConfig.String("dbType")
	dbHost := beego.AppConfig.String("dbHost")
	dbName := beego.AppConfig.String("dbName")
	dbUrl := beego.AppConfig.String("dbUrl")
	dbPasaswd := beego.AppConfig.String("dbPasaswd")
	dbUser := beego.AppConfig.String("dbUser")
	connect := dbUser + ":" + dbPasaswd + "@tcp(" + dbUrl + ":" + dbHost + ")/" + dbName + "?charset=utf8"

	orm.RegisterDataBase("default", dbType, connect, 30)
	// register model
	orm.RegisterModel(new(User))

	// create table
	//orm.RunSyncdb("default", false, true)
}
