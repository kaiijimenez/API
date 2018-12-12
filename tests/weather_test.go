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
	_ "github.com/kaiijimenez/API/routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../../"+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestGet is a sample to run an endpoint test
func TestWeatherApi(t *testing.T) {
	var r libraries.Response
	var e libraries.ErrorResponse
	valid := fmt.Sprintf(beego.AppConfig.String("lhost"), "Mexico", "mx")
	empty := fmt.Sprintf(beego.AppConfig.String("lhost"), "", "")
	mismatch := fmt.Sprintf(beego.AppConfig.String("lhost"), "Bogota", "tf")
	tables := []struct {
		url      string
		response interface{}
		mlogs    string
	}{
		{valid, r, "Test Getting Response from Weather Endpoint"},
		{empty, e, "Testing when city or country or both is not sending in the request"},
		{mismatch, e, "Testing when city or country are mismatching in the request"},
	}

	for _, table := range tables {
		req, _ := http.NewRequest("GET", table.url, nil)
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, req)
		err := json.Unmarshal([]byte(w.Body.String()), &table.response)
		Convey(table.mlogs, t, func() {
			Convey("Status should be 200", func() {
				So(w.Code, ShouldEqual, 200)
			})
			Convey("Body must be a valid JSON", func() {
				So(err == nil, ShouldBeTrue)
			})
			Convey("Body must contain all fields", func() {
				So(
					table.response != "",
					ShouldBeTrue,
				)
			})
		})
	}
}
