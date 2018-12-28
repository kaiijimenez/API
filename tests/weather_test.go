package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/kaiijimenez/API/libraries"
	"github.com/kaiijimenez/API/libraries/structs"

	_ "github.com/kaiijimenez/API/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../../"+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp/weatherapidb?charset=utf8")
	o := orm.NewOrm()
	orm.Debug = true
	o.Using("default")
}

// TestGet is a sample to run an endpoint test
func TestWeatherApiSuccess(t *testing.T) {
	var r structs.Response
	valid := fmt.Sprintf(beego.AppConfig.String("localhost"), "Paris", "fr")

	req, _ := http.NewRequest("GET", valid, nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)
	err := json.Unmarshal([]byte(w.Body.String()), &r)

	Convey("Test getting info from Mexico, mx", t, func() {
		Convey("Status should be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Body must be a Valid Json response", func() {
			So(err == nil, ShouldBeTrue)
		})
		Convey("Response should contain all data", func() {
			So(
				r.Location_name != "" && r.Temperature != "" &&
					r.Wind != "" && r.Pressure != "" && r.Humidity != "" &&
					r.Geo_coordinates != "" && r.Sunrise != "" && r.Sunset != "" &&
					r.Requested_time != "", ShouldBeTrue,
			)
		})
	})

}

func TestWeatherApiErrors(t *testing.T) {
	var e structs.ErrorResponse
	empty := fmt.Sprintf(beego.AppConfig.String("localhost"), "", "")
	mismatch := fmt.Sprintf(beego.AppConfig.String("localhost"), "Bogota", "tf")

	tables := []struct {
		url   string
		mlogs string
	}{
		{empty, "Testing when city or country or both are not sending in the request"},
		{mismatch, "Testing when city or country are mismatching in the request"},
	}

	for _, table := range tables {
		req, _ := http.NewRequest("GET", table.url, nil)
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, req)
		err := json.Unmarshal([]byte(w.Body.String()), &e)

		Convey(table.mlogs, t, func() {
			Convey("Status should be 200", func() {
				So(w.Code, ShouldEqual, 200)
			})
			Convey("Body must be a valid JSON", func() {
				So(err, ShouldBeNil)
			})
			Convey("Body must code and message of error", func() {
				So(
					e.Code != "" && e.Message != "",
					ShouldBeTrue,
				)
			})
		})
	}
}


//Unit testing of functions in libraries file
func TestGetJsonResponse(t *testing.T) {
	endapi := beego.AppConfig.String("weatherprovider")
	files := beego.AppConfig.String("fileprovider")
	tables := []struct {
		prov    string
		city    string
		country string
		mlogs   string
	}{
		{endapi, "bogota", "co", "Getting response from endpoint"},
		{files, "mexico", "mx", "Getting response from file"},
	}
	for _, table := range tables {
		jres := libraries.GetJsonResponse(table.prov, table.city, table.country)
		Convey(table.mlogs, t, func() {
			Convey("Response from GetJsonResponse should not return a nil value", func() {
				So(jres, ShouldNotBeNil)
			})
		})
	}
}

func TestReadJsonFile(t *testing.T) {
	jres := libraries.ReadJsonFile("argentina", "us")
	Convey("Getting response from file", t, func() {
		Convey("When a file doesnt exist the response should be nil", func() {
			So(jres, ShouldBeNil)
		})
	})
}

