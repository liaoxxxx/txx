package main

import (
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	_ "time"
	_ "untitle_go_project1/routers"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", true, false)
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
	u := new(User)
	u.Password = "liao993501756"
	u.Name = "liaoxxxx"
	u.Status = 1
	u.Is_delete = 1
	u.Created_at = time.Now().Unix()
	u.Update_at = time.Now().Unix()
	o.Insert(u)
	return
}
