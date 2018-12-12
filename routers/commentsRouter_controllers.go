package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/kaiijimenez/API/controllers:WeatherController"] = append(beego.GlobalControllerRouter["github.com/kaiijimenez/API/controllers:WeatherController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/r`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
