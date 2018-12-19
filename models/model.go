package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	_"github.com/astaxie/beego"

)

type User struct {
	Id int
	User_name string
	Salt string
	Passwd string
	Status byte
	Nick_name string

}

func init()  {

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterModel(new(User))
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.RunSyncdb("default", false, true)

}
