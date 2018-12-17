package main

import (
	"fmt"

	_ "github.com/kaiijimenez/API/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.Debug = true
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
	logs.Info("App started")
}
