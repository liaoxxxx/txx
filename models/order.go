package models

import (
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Order struct {
	Id        int
	User_name string
	Salt      string
	Passwd    string
	Status    byte
	Nick_name string
}

func init() {
	orm.RegisterModel(new(Order))
}
