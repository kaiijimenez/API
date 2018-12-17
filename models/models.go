package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Weather struct {
	Id            int
	LocationName  string
	Temperature   string
	Wind          string
	Cloudiness    string
	Pressure      string
	Humidity      string
	Sunrise       time.Time
	Sunset        time.Time
	Lon           float64
	Lat           float64
	RequestedTime time.Time
	Timestamp     time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Weather))
}

func (this *Weather) Table() string {
	return "weather"
}
