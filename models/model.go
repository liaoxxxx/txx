package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8", 30)
	// register model
	//orm.RegisterModel(new(User))

	// create table
	orm.RunSyncdb("default", false, true)
}
