package model

type Data struct {
	Resp []Temprature `json:"data"`
}
type Temprature struct {
	Date    string  `json:"date"`
	TempAvg float64 `json:"tavg"`
	TempMin float64 `json:"tmin"`
	TempMax float64 `json:"tmax"`
}
