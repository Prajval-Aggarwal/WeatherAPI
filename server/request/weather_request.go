package request

type WeatherRequest struct {
	City      []string `json:"cities"`
	StartDate string   `json:"start"`
	EndDate   string   `json:"end"`
}
