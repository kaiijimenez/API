package controllers

import (
	"fmt"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/kaiijimenez/API/libraries"
	"github.com/astaxie/beego"
)

type WeatherController struct {
	beego.Controller
}

var (
	wg         sync.WaitGroup
	workernums = 5
)


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

	results := make(chan interface{}, 10)
	for i := 0; i < workernums; i++ {
		time.Sleep(time.Second * 10)
		wg.Add(1)
		go work(i, results, city, country)
	}
	logs.Info("Showing response to client")
	wc.Data["json"] = <-results
	wc.ServeJSON()
	wg.Wait()

}

func work(id int, result chan interface{}, city, country string) {
	fmt.Println("Worker : ", id)
	fmt.Println("Getting response from endpoint")
	gotRest := libraries.GetResponse(city, country)
	result <- gotRest
	wg.Done()
}
