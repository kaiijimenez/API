package controllers

import (
	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/kaiijimenez/API/tasks"
)

type ScheduleController struct {
	beego.Controller
}

type Success struct {
	city    string
	country string
}

//@Title PUT Scheduler Perform a regular check (Every 1 hour) and persist it in the DB
//@Description Persist city in the db
//@Param city query string false "name of the City Example: Bogota"
//@Param country query string false "Country is a country code of two characters in lowercase. Example: co"
//@Success 202 {"city": city, "country":country}
//@Accept json
//@router /r [put]
func (sc *ScheduleController) Put() {
	city := sc.GetString("city")
	country := sc.GetString("country")
	logs.Info("Put controller scheduler")
	tasks.WeatherTask(city, country)
	sc.Ctx.ResponseWriter.WriteHeader(202)
	sc.Ctx.Output.Body([]byte(city))
	return
}
