package structs

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

type NotFoundRespose struct {
	Code    string `json:"cod"`
	Message string `json:"message"`
}
