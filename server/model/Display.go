package model

type Data struct {
	Cityname string   `json:"cityName"`
	Info     CityInfo `json:"cityInfo"`
}

type GroupData struct {
	CityName string     `json:"cityName"`
	Info     TempStruct `json:"cityInfo"`
}

type TempStruct struct {
	Data []TempData `json:"data"`
}

type TempData struct {
	Name    string
	TempAvg float64
}
