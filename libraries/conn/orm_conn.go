package conn

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/kaiijimenez/API/libraries/structs"
	"github.com/kaiijimenez/API/models"
)

func Inserting(response structs.Response, jresponse *structs.JsonResponse, sunrise, sunset, rtime time.Time) (*structs.Response, error) {
	logs.Info("Connecting with DB")
	//Saving data

	o := orm.NewOrm()
	o.Using("default")

	db := new(models.Weather)
	queryseter := o.QueryTable("weather").Filter("location_name", response.Location_name).OrderBy("-id").Limit(1)
	qerr := queryseter.One(db)

	if qerr == orm.ErrNoRows {
		logs.Info("Inserting new values")
		db.LocationName = response.Location_name
		db.Temperature = response.Temperature
		db.Wind = response.Wind
		db.Cloudiness = jresponse.Weather[0].Description
		db.Pressure = response.Pressure
		db.Humidity = response.Humidity
		db.Lat = jresponse.Coord.Lat
		db.Lon = jresponse.Coord.Lon
		db.Sunrise = sunrise
		db.Sunset = sunset
		db.RequestedTime = rtime

		_, err := o.Insert(db)
		if err != nil {
			return nil, err
		}
	} else if time.Now().Sub(db.Timestamp).Seconds() > 300 {
		logs.Info("Inserting new values if timestamp is > 300")
		newcol := new(models.Weather)
		newcol.LocationName = response.Location_name
		newcol.Temperature = response.Temperature
		newcol.Wind = response.Wind
		newcol.Cloudiness = jresponse.Weather[0].Description
		newcol.Pressure = response.Pressure
		newcol.Humidity = response.Humidity
		newcol.Lat = jresponse.Coord.Lat
		newcol.Lon = jresponse.Coord.Lon
		newcol.Sunrise = sunrise
		newcol.Sunset = sunset
		newcol.RequestedTime = rtime

		_, err := o.Insert(newcol)
		if err != nil {
			return nil, err
		}
	} else {
		logs.Info("Returning values from db")
		response.Location_name = db.LocationName
		response.Temperature = db.Temperature
		response.Wind = db.Wind
		response.Cloudiness = db.Cloudiness
		response.Pressure = fmt.Sprintf("%.0f hpa", jresponse.Main.Pressure)
		response.Humidity = db.Humidity
		response.Sunrise = fmt.Sprintf("%02d:%02d", db.Sunrise.Hour(), db.Sunrise.Minute())
		response.Sunset = fmt.Sprintf("%02d:%02d", db.Sunset.Hour(), db.Sunset.Minute())
		response.Geo_coordinates = fmt.Sprintf("%v", []float64{db.Lat, db.Lon})
		response.Requested_time = fmt.Sprintf("%v", db.RequestedTime)
	}
	return &response, nil
}
