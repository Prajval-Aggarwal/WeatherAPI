package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"main/server/model"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func GetCoordinates(params url.Values) (*[]model.Coordinates, error) {
	//params := c.Request.URL.Query()
	fmt.Println("params is", params)
	var data []model.Coordinates
	for _, v := range params {

		split := strings.Split(v[0], ",")
		city := split[0]
		state := split[1]
		fmt.Println("state is:", state)

		apiUrl := fmt.Sprintf("https://api.openweathermap.org/geo/1.0/direct?q=%s&limit=10&appid=%s", city, os.Getenv("OPEN_WEATHER_API"))
		resp, err := http.Get(apiUrl)
		if err != nil {
			return nil, errors.New("failed to make API call")
		}
		var cord []model.Coordinates
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New("failed to read response body")
		}
		err = json.Unmarshal(body, &cord)
		if err != nil {
			fmt.Println("error is:", err)
			return nil, err
		}
		for _, val := range cord {
			if state == val.State {
				data = append(data, val)
				break
			}
		}

	}
	return &data, nil
}
