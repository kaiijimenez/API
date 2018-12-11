package libraries

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

type JsonResponse struct {
	Coord   CoordStruct     `json:"coord"`
	Weather []WeatherStruct `json:"weather"`
	Main    MainStruct      `json:"main"`
	Wind    WindStruct      `json:"wind"`
	Sys     SysStruct       `json:"sys"`
	Rtime   int64           `json:"dt"`
	Name    string          `json:"name"`
}

type CoordStruct struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type WeatherStruct struct {
	Description string `json:"description"`
}

type MainStruct struct {
	Temperature float64 `json:"temp"`
	Pressure    float64 `json:"pressure"`
	Humidity    float64 `json:"humidity"`
}

type WindStruct struct {
	Speed    float64 `json:"speed"`
	Degrates float64 `json:"deg"`
}

type SysStruct struct {
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type Response struct {
	Location_name   string
	Temperature     string
	Wind            string
	Cloudiness      string
	Pressure        string
	Humidity        string
	Sunrise         string
	Sunset          string
	Geo_coordinates string
	Requested_time  string
}

type ErrorResponse struct {
	Code    string
	Message string
}

func GetResponse(city, country string) interface{} {
	var response Response
	var eresponse ErrorResponse
	var jresponse JsonResponse
	weather := GetConfig("weather")
	base := GetConfig("base_url")
	appid := GetConfig("appid")

	// checks whether the city or the country are empty or if the Country/City are not compatible
	if city == "" || country == "" {
		logs.Critical("City or country not found or empty")
		eresponse.Code = "404"
		eresponse.Message = "City or country not found or empty!"
		return eresponse
	}
	we := fmt.Sprintf(weather, city, country)
	uri := fmt.Sprintf("%s%s%s", base, we, appid)
	fmt.Println(base)
	fmt.Println(uri)
	//uri := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%v,%v&appid=1508a9a4840a5574c822d70ca2132032", city, country)
	res := httplib.Get(uri) //get the uri sending
	str, e := res.String()  //return raw response body
	CheckErrors("Error in the raw response: ", e)
	err := json.Unmarshal([]byte(str), &jresponse)
	CheckErrors("Error trying to unmarshal the data: ", err)

	//getting time
	sunrise := time.Unix(jresponse.Sys.Sunrise, 0)
	sunset := time.Unix(jresponse.Sys.Sunset, 0)
	rtime := time.Unix(jresponse.Rtime, 0)
	//getting wind values
	b, w, wd := beaufort(jresponse.Wind.Speed), jresponse.Wind.Speed, windir(jresponse.Wind.Degrates)

	response.Location_name = fmt.Sprintf("%v, %v", city, strings.ToUpper(country))
	response.Temperature = getTemperature(jresponse.Main.Temperature)
	response.Wind = fmt.Sprintf("%s, %.2f m/s, %s", b, w, wd)
	response.Cloudiness = fmt.Sprintf("%s", jresponse.Weather[0].Description)
	response.Pressure = fmt.Sprintf("%.0f hpa", jresponse.Main.Pressure)
	response.Humidity = fmt.Sprintf("%.0f%%", jresponse.Main.Humidity)
	response.Sunrise = fmt.Sprintf("%02d:%02d", sunrise.Hour(), sunrise.Minute())
	response.Sunset = fmt.Sprintf("%02d:%02d", sunset.Hour(), sunset.Minute())
	response.Geo_coordinates = fmt.Sprintf("%v", []float64{jresponse.Coord.Lat, jresponse.Coord.Lon})
	response.Requested_time = fmt.Sprintf("%v", rtime)
	return response
}

func CheckErrors(s string, e error) {
	if e != nil {
		logs.Critical(s, e)
	}
}

//Gets wind speed
func beaufort(speed float64) string {
	var b string
	if speed < 0.3 {
		b = "Calm"
	} else if 0.3 <= speed && speed <= 1.5 {
		b = "Light Air"
	} else if 1.6 <= speed && speed <= 3.3 {
		b = "Light Breeze"
	} else if 3.4 <= speed && speed <= 5.5 {
		b = "Gentle Breeze"
	} else if 5.5 <= speed && speed <= 7.9 {
		b = "Moderate Breeze"
	} else if 8.0 <= speed && speed <= 10.7 {
		b = "Fresh Breeze"
	} else if 10.8 <= speed && speed <= 13.8 {
		b = "Strong Breeze"
	} else if 13.9 <= speed && speed <= 17.1 {
		b = "Near Gale"
	} else if 17.2 <= speed && speed <= 20.7 {
		b = "Gale"
	} else if 20.8 <= speed && speed <= 24.4 {
		b = "Strong Gale"
	} else if 24.5 <= speed && speed <= 28.4 {
		b = "Storm"
	} else if 28.5 <= speed && speed <= 32.6 {
		b = "Violent Storm"
	} else if speed >= 32.7 {
		b = "Hurricane Force"
	}
	return b
}

//Gets wind direction
func windir(deg float64) string {
	var d string
	if 348.75 <= deg && deg <= 11.25 {
		d = "North"
	} else if 11.26 <= deg && deg <= 33.75 {
		d = "North-NorthEast"
	} else if 33.76 <= deg && deg <= 56.25 {
		d = "NorthEast"
	} else if 56.26 <= deg && deg <= 78.75 {
		d = "East-NorthEast"
	} else if 78.76 <= deg && deg <= 101.25 {
		d = "East"
	} else if 101.26 <= deg && deg <= 123.75 {
		d = "East-SouthEast"
	} else if 123.26 <= deg && deg <= 146.75 {
		d = "SouthEast"
	} else if 146.26 <= deg && deg <= 168.75 {
		d = "South-SouthEast"
	} else if 168.26 <= deg && deg <= 191.75 {
		d = "South"
	} else if 191.26 <= deg && deg <= 213.75 {
		d = "South-SouthWest"
	} else if 213.26 <= deg && deg <= 236.75 {
		d = "SouthWest"
	} else if 236.26 <= deg && deg <= 258.75 {
		d = "West-SouthWest"
	} else if 258.26 <= deg && deg <= 281.75 {
		d = "West"
	} else if 281.26 <= deg && deg <= 303.75 {
		d = "West-NorthWest"
	} else if 303.26 <= deg && deg <= 326.75 {
		d = "NorthWest"
	} else if 326.26 <= deg && deg <= 348.75 {
		d = "North-NorthWest"
	}
	return d
}

var getTemperature = func(kelvin float64) string {
	return fmt.Sprintf("%.0f °C", kelvin-273.15)
}

func GetConfig(s string) string {
	return beego.AppConfig.String(s)
}
