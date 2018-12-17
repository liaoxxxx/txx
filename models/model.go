package models

import "github.com/astaxie/beego/orm"

type User struct {
	id int
	manager_name string
	salt string
	passwd string
	status byte

}

func init()  {
	orm.RegisterDataBase("default","mysql","")
}
