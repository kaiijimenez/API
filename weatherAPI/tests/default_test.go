package test

import (
	_ "API/weatherAPI/routers"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestGet is a sample to run an endpoint test
func TestWeatherApi(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/weather/r?city=Mexico&country=mx", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	//beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

	var jsonResponse map[string]interface{}

	err := json.Unmarshal([]byte(w.Body.String()), &jsonResponse)
	fmt.Println(jsonResponse)

	Convey("Test Getting Response from Weather Endpoint", t, func() {
		Convey("Status should be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Body must be a valid JSON", func() {
			So(err == nil, ShouldBeTrue)
		})
		Convey("Body must contain all fields", func() {
			So(
				jsonResponse["location_name"] != nil &&
					jsonResponse["temperature"] != nil &&
					jsonResponse["wind"] != nil &&
					jsonResponse["pressure"] != nil &&
					jsonResponse["humidity"] != nil &&
					jsonResponse["geo_coordinates"] != nil &&
					jsonResponse["sunrise"] != nil &&
					jsonResponse["sunset"] != nil &&
					jsonResponse["requested_time"] != nil,
				ShouldBeTrue,
			)
		})
	})
}
