package model

type Coordinates struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	State     string  `json:"state"`
}
