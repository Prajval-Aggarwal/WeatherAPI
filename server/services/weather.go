package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"main/server/model"
	"main/server/response"
	"main/server/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetCoordinates(params url.Values) (*[]model.ResponseCoordinate, error) {
	//params := c.Request.URL.Query()
	fmt.Println("params is", params)
	var coordinates []model.ResponseCoordinate
	for key, v := range params {

		if key == "start" || key == "end" {
			continue
		}
		split := strings.Split(v[0], ",")
		city := split[0]
		state := split[1]
		fmt.Println("city is", city)
		fmt.Println("state is:", state)

		apiUrl := fmt.Sprintf("https://trueway-geocoding.p.rapidapi.com/Geocode?address=%s", city)
		req, err := http.NewRequest("GET", apiUrl, nil)
		if err != nil {
			return nil, errors.New("failed to make API call")
		}
		req.Header.Add("X-RapidAPI-Key", "877d7321bdmsh5db7b7a54b66d8fp168429jsn4835e587d80e")
		req.Header.Add("X-RapidAPI-Host", "trueway-geocoding.p.rapidapi.com")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		// Read the response body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, errors.New("failed to read response body")
		}
		var singleCord model.Coordinate
		err = json.Unmarshal(body, &singleCord)
		if err != nil {
			fmt.Println("error is:", err)
			return nil, err
		}
		tempResponse := model.ResponseCoordinate{
			Name:      singleCord.Res[0].Address,
			Latitude:  singleCord.Res[0].Location.Latitude,
			Longitude: singleCord.Res[0].Location.Longitude,
		}
		coordinates = append(coordinates, tempResponse)

	}
	return &coordinates, nil
}

func ExtractingData(Latitude float64, Longitude float64, startDate string, endDate string, Period string) (*model.CityInfo, error) {
	var data model.CityInfo
	apiUrl := fmt.Sprintf("https://meteostat.p.rapidapi.com/point/%s?lat=%v&lon=%v&start=%v&end=%v", Period, Latitude, Longitude, startDate, endDate)

	fmt.Println("apiUrl", apiUrl)
	req, _ := http.NewRequest("GET", apiUrl, nil)

	req.Header.Add("X-RapidAPI-Key", "877d7321bdmsh5db7b7a54b66d8fp168429jsn4835e587d80e")
	req.Header.Add("X-RapidAPI-Host", "meteostat.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil

}

func Daily(ctx *gin.Context) {
	//get data from params
	params := ctx.Request.URL.Query()
	fmt.Println("params is:", params)
	var totalData []model.Data
	//get start and end date from params'

	startDate := params.Get("start")
	endDate := params.Get("end")
	//first get the coordinates
	coordinates, err := GetCoordinates(params)

	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	fmt.Println("coordinates are:", coordinates)
	for _, coordinate := range *coordinates {
		data, err := ExtractingData(coordinate.Latitude, coordinate.Longitude, startDate, endDate, "daily")
		if err != nil {
			response.ErrorResponse(ctx, 400, err.Error())
			return
		}
		tempData := model.Data{
			Cityname: coordinate.Name,
			Info:     *data,
		}
		totalData = append(totalData, tempData)
	}
	response.ShowResponse("Success", 200, "Data fetched successfully", totalData, ctx)
}

func Weekly(ctx *gin.Context) {
	//get data from params
	params := ctx.Request.URL.Query()
	fmt.Println("params is:", params)
	//get start and end date from params'
	var DataSlice []model.GroupData
	startDate := params.Get("start")
	endDate := params.Get("end")
	//first get the coordinates
	coordinates, err := GetCoordinates(params)

	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	for _, coordinate := range *coordinates {
		Data, err := ExtractingData(coordinate.Latitude, coordinate.Longitude, startDate, endDate, "daily")
		if err != nil {
			response.ErrorResponse(ctx, 400, err.Error())
			return
		}
		var tempSLice []model.TempData
		counter := 1
		fmt.Println("length of data is", len(Data.Resp)/7)
		for i := 0; i < len(Data.Resp); i += 7 {
			tempAvg := 0.0
			tempAdd := 0.0

			for j := i; j < i+7 && j < len(Data.Resp); j++ {
				tempAdd += Data.Resp[j].TempAvg
			}
			tempAvg = tempAdd / 7
			temp := model.TempData{
				Name:    "Week" + strconv.Itoa(counter),
				TempAvg: utils.RoundFloat(tempAvg, 3),
			}
			tempSLice = append(tempSLice, temp)
			counter++
		}
		monthly := model.GroupData{
			CityName: coordinate.Name,
		}
		monthly.Info.Data = tempSLice
		DataSlice = append(DataSlice, monthly)

	}
	response.ShowResponse("Success", 200, "Data fetched successfully", DataSlice, ctx)
}

func Monthly(ctx *gin.Context) {
	params := ctx.Request.URL.Query()
	fmt.Println("params is:", params)
	var totalData []model.Data
	//get start and end date from params'

	startDate := params.Get("start")
	endDate := params.Get("end")
	//first get the coordinates
	coordinates, err := GetCoordinates(params)
	fmt.Println("coordinates is", coordinates)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	for _, coordinate := range *coordinates {
		data, err := ExtractingData(coordinate.Latitude, coordinate.Longitude, startDate, endDate, "monthly")
		if err != nil {
			response.ErrorResponse(ctx, 400, err.Error())
			return
		}
		tempData := model.Data{
			Cityname: coordinate.Name,
			Info:     *data,
		}
		totalData = append(totalData, tempData)

	}
	response.ShowResponse("Success", 200, "Data fetched successfully", totalData, ctx)
}

func Yearly(ctx *gin.Context) {
	//get data from params
	params := ctx.Request.URL.Query()
	fmt.Println("params is:", params)
	//get start and end date from params'

	var DataSlice []model.GroupData
	startDate := params.Get("start")
	endDate := params.Get("end")
	//first get the coordinates
	coordinates, err := GetCoordinates(params)

	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	for _, coordinate := range *coordinates {
		Data, err := ExtractingData(coordinate.Latitude, coordinate.Longitude, startDate, endDate, "monthly")
		if err != nil {
			response.ErrorResponse(ctx, 400, err.Error())
			return
		}
		var tempSLice []model.TempData
		counter := 1
		fmt.Println("length of data is", len(Data.Resp))
		for i := 0; i < len(Data.Resp); i += 12 {
			tempAvg := 0.0
			tempAdd := 0.0
			for j := i; j < i+12 && j < len(Data.Resp); j++ {
				tempAdd += Data.Resp[j].TempAvg
			}
			tempAvg = tempAdd / 12

			temp := model.TempData{
				Name:    "Year " + strconv.Itoa(counter),
				TempAvg: utils.RoundFloat(tempAvg, 3),
			}
			tempSLice = append(tempSLice, temp)
			counter++

		}
		yearly := model.GroupData{
			CityName: coordinate.Name,
		}
		yearly.Info.Data = tempSLice

		DataSlice = append(DataSlice, yearly)

	}
	response.ShowResponse("Success", 200, "Data fetched successfully", DataSlice, ctx)
}
