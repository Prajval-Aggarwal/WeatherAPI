package model

type Coordinate struct {
	Res []struct {
		Address  string `json:"address"`
		Location struct {
			Latitude  float64 `json:"lat"`
			Longitude float64 `json:"lng"`
		} `json:"location"`
	} `json:"results"`
}

type ResponseCoordinate struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}
