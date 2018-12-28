package tasks

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
	"github.com/kaiijimenez/API/libraries"
)

func WeatherTask(city, country string) error {
	logs.Info("Weather task")
	weathertask := toolbox.NewTask("weathertask", "* * */1 * * *", func() error {
		libraries.GetResponse(city, country)
		return nil
	})
	toolbox.AddTask("weathertask", weathertask)
	toolbox.StartTask()
	return nil
}
