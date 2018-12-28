package libraries

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/kaiijimenez/API/libraries/conn"
	"github.com/kaiijimenez/API/libraries/structs"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
)

var (
	response  structs.Response
	jresponse structs.JsonResponse
	notfound  structs.NotFoundRespose
)

func GetResponse(city, country string) interface{} {
	logs.Info("Retrieving weather info")
	//provider file or openapi config variables:
	//weatherprovider and fileprovider
	prov := beego.AppConfig.String("weatherprovider")
	jresponse := GetJsonResponse(prov, city, country)

	if jresponse == nil {
		return EResponse()
	}

	//getting time
	sunrise := time.Unix(jresponse.Sys.Sunrise, 0)
	sunset := time.Unix(jresponse.Sys.Sunset, 0)
	rtime := time.Unix(jresponse.Rtime, 0)
	//getting wind values
	b, w, wd := beaufort(jresponse.Wind.Speed), jresponse.Wind.Speed, windir(jresponse.Wind.Degrates)

	//getting the response with human redable data
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

	//Inserting into DB returns the response retrieve from db (in case they exist) to return to client
	//Return err in case there is an error trying to insert in the db
	resp, err := conn.Inserting(response, jresponse, sunrise, sunset, rtime)
	if resp != nil && err == nil {
		return resp
	} else if resp == nil && err != nil {
		return InsertError()
	}
	return response
}

func CheckErrors(s string, e error) {
	if e != nil {
		logs.Critical(s, e)
	}
}

func EResponse() structs.ErrorResponse {
	var eresponse structs.ErrorResponse
	logs.Critical("City or country not found or empty")
	eresponse.Code = "404"
	eresponse.Message = "City not found"
	return eresponse
}


func InsertError() structs.ErrorResponse {
	var eresponse structs.ErrorResponse
	logs.Critical("Error trying to insert in DB")
	eresponse.Code = "500"
	eresponse.Message = "Error trying to insert in DB"
	return eresponse
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


func ReadJsonFile(city, country string) *structs.JsonResponse {
	//Reading data from file
	filename := fmt.Sprintf("%s_%s.json", strings.ToLower(city), strings.ToLower(country))
	path := filepath.Join("files/", filename)

	readfile, err := ioutil.ReadFile(path)
	if err != nil {
		CheckErrors("Error trying to read the file", err)
		return nil
	}
	er := json.Unmarshal(readfile, &jresponse)
	if er != nil {
		CheckErrors("Error trying to unmarshal the data from file: ", er)
		return nil
	}
	return &jresponse
}

func GetData(city, country string) string {
	//getting uri
	base := beego.AppConfig.String("base_url")
	appid := beego.AppConfig.String("appid")
	uri := fmt.Sprintf(base, city, country, appid)
	res := httplib.Get(uri)
	str, e := res.String()
	CheckErrors("Error in the raw response: ", e)
	return str
}

func GetJsonResponse(prov, city, country string) *structs.JsonResponse {
	if prov == "api.openweathermap.org" {
		//Getting data from Endpoint:
		str := GetData(city, country)
		err := json.Unmarshal([]byte(str), &notfound)
		if err != nil {
			er := json.Unmarshal([]byte(str), &jresponse)
			CheckErrors("Error trying to unmarshall from endpoint: ", er)
			return &jresponse
		}
		return nil
	}
	//SaveJsonFile(city, country)
	//Get data from File
	return ReadJsonFile(city, country)
}
