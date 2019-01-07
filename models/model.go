package models

import (
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id        int
	User_name string
	Salt      string
	Passwd    string
	Status    byte
	Nick_name string
}

func init() {
	orm.RegisterModel(new(User))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
}
func (o Order) test() {

}
