// @APIVersion 1.0.0
// @Title WeatherAPI
// @Weather API
// @Contact karina.jimenez@globant.com
package routers

import (
	"github.com/kaiijimenez/API/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/weather",
			beego.NSInclude(&controllers.WeatherController{}),
		),
	)
	beego.AddNamespace(ns)
}
