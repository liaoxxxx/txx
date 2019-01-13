package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "untitle_go_project1/routers"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8")
	orm.RegisterModel(new(User))
	//orm.RunSyncdb("default", true, false)
}

type User struct {
	Id         uint32
	Name       string
	Password   string
	Salt       string
	Status     byte
	Is_delete  byte
	Created_at int64
	Update_at  int64
}

func main() {
	o := orm.NewOrm()
	o.Using("default")
	u := new(User)
	u.Name = "liaoxx124"
	err := o.Read(u, "Name")
	if err != nil {
		beego.Info("查询失败", err)
	} else {
		fmt.Println("查询成功", u.Status)
	}
	return
}
