package controllers

import (
	"github.com/astaxie/beego/logs"

	"github.com/kaiijimenez/API/libraries"

	"github.com/astaxie/beego"
)

type WeatherController struct {
	beego.Controller
}

//@Title Get Json Response
//@Description Get the Response from the endpoint of weather api
//@Param city query string false "name of the City Example: Bogota"
//@Param country query string false "Country is a country code of two characters in lowercase. Example: co"
//@Success 200 {object} "Success"
//@Failure 400 Bad Request
//@Failure 400 Not Found
//@Accept json
//@router /r [get]
func (wc *WeatherController) Get() {
	city := wc.GetString("city")
	country := wc.GetString("country")
	//worker pool
	logs.Info("Getting the response from the endpoint")
	wc.Data["json"] = libraries.GetResponse(city, country)
	wc.ServeJSON()

}
