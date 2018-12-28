package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/kaiijimenez/API/routers"
	_ "github.com/kaiijimenez/API/tasks"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// "$USER:PASS@tcp($HOST:$PORT)/DBNAME",
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(weatherdb)/weatherapidb?charset=utf8")
}

func main() {

	fmt.Println("Running!")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
