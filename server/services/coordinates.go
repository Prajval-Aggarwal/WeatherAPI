package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"main/server/model"
	"main/server/response"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetCoordinates(params url.Values) (*[]model.Coordinates, error) {
	//params := c.Request.URL.Query()
	fmt.Println("params is", params)
	var data []model.Coordinates
	for key, v := range params {

		if key == "start" || key == "end" {
			continue
		}
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

func Daily(ctx *gin.Context) {
	//get data from params
	params := ctx.Request.URL.Query()
	fmt.Println("params is:", params)
	//get start and end date from params'

	startDate := params.Get("start")
	sDate, _ := time.Parse("2006-01-02", startDate)
	ssDate := sDate.Truncate(time.Hour)
	endDate := params.Get("end")

	eDate, _ := time.Parse("2006-01-02", endDate)
	eeDate := eDate.Truncate(time.Hour)

	//first get the coordinates
	coordinates, err := GetCoordinates(params)

	fmt.Println("coordinates is:", coordinates)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	for _, coordinate := range *coordinates {
		fmt.Println("C", coordinate)
		//var data model.Data
		apiUrl := fmt.Sprintf("https://meteostat.p.rapidapi.com/point/daily?lat=%v&lon=%v&start=%v&end=%v", coordinate.Latitude, coordinate.Longitude, ssDate, eeDate)
		fmt.Println("apiUrl", apiUrl)
		req, _ := http.NewRequest("GET", apiUrl, nil)
		//https://meteostat.p.rapidapi.com/point/daily?lat=30.7046&lon=76.7179&start=2011-01-01&end=2020-01-31
		//https://meteostat.p.rapidapi.com/point/daily?lat=30.9090157&lon=75.851601&start=2001-01-01&end=2020-01-01
		//https://meteostat.p.rapidapi.com/point/daily?lat=30.9090&lon=75.8516&start=2011-01-01&end=2020-01-31

		req.Header.Add("X-RapidAPI-Key", "877d7321bdmsh5db7b7a54b66d8fp168429jsn4835e587d80e")
		req.Header.Add("X-RapidAPI-Host", "meteostat.p.rapidapi.com")
		fmt.Println("headder are", req.Header)
		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		ctx.Data(200, "application/json", body)
		// err = json.Unmarshal(body, &data)
		// if err != nil {
		// 	fmt.Println("error is:", err)
		// 	return
		// }
		// ctx.JSON(200, data)
	}

}
