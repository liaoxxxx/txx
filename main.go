package main

import (
	"github.com/astaxie/beego"
	_ "untitle_go_project1/models"
	_ "untitle_go_project1/routers"
)

func main() {
	beego.Run()
}
