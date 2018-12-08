package libraries

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

func GetResponse(city, country string) map[string]interface{} {
	var jsonResponse map[string]interface{}
	response := make(map[string]interface{})

	uri := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%v,%v&appid=1508a9a4840a5574c822d70ca2132032", city, country)
	res := httplib.Get(uri) //get the uri sending
	res.Debug(true)
	str, e := res.String() //return raw response body
	CheckErrors("Error returning the raw response: ", e)
	err := json.Unmarshal([]byte(str), &jsonResponse)
	CheckErrors("Error trying to unmarshal the data: ", err)

	//getting the json data
	coord := jsonResponse["coord"].(map[string]interface{})
	wind := jsonResponse["wind"].(map[string]interface{})
	main := jsonResponse["main"].(map[string]interface{})
	sys := jsonResponse["sys"].(map[string]interface{})
	clouds := jsonResponse["clouds"].(map[string]interface{})
	//getting time
	sunrise := time.Unix(int64(sys["sunrise"].(float64)), 0)
	sunset := time.Unix(int64(sys["sunset"].(float64)), 0)
	timer := time.Unix(int64(jsonResponse["dt"].(float64)), 0)
	//getting wind values
	b, w, wd := beaufort(wind["speed"].(float64)), wind["speed"].(float64), windir(wind["deg"].(float64))

	response["location_name"] = fmt.Sprintf("%s, %s", jsonResponse["name"], sys["country"])

	response["temperature"] = fmt.Sprintf("%.0f °C", main["temp"].(float64)-273.15)

	response["wind"] = fmt.Sprintf("%s, %.2f m/s, %s", b, w, wd)

	response["cloudiness"] = fmt.Sprintf("%s", cloudiness(clouds["all"].(float64)))

	response["pressure"] = fmt.Sprintf("%.0f hpa", main["pressure"])

	response["humidity"] = fmt.Sprintf("%.0f%%", main["humidity"])

	response["sunrise"] = fmt.Sprintf("%02d:%02d", sunrise.Hour(), sunrise.Minute())

	response["sunset"] = fmt.Sprintf("%02d:%02d", sunset.Hour(), sunset.Minute())

	response["geo_coordinates"] = []float64{coord["lon"].(float64), coord["lat"].(float64)}

	response["requested_time"] = fmt.Sprintf(
		"%v",
		time.Date(timer.Year(), timer.Month(), timer.Day(), timer.Hour(), timer.Minute(), timer.Second(), timer.Nanosecond(), timer.Location()))
	return response
}

func CheckErrors(s string, e error) {
	if e != nil {
		logs.Error(s, e)
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

//cloudiness
func cloudiness(all float64) string {
	var c string
	if 0.0 <= all && all <= 25.0 {
		c = "Clear Skies"
	} else if 26.0 <= all && all <= 75.0 {
		c = "Scattered Clouds"
	} else if 76.0 <= all && all <= 99.0 {
		c = "Broken Clouds"
	} else if all == 100.0 {
		c = "Overcast"
	}
	return c
}
