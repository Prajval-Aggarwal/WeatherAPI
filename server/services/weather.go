package services

import (
	"fmt"
	"log"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/utils"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var wg = sync.WaitGroup{}
var mut = sync.Mutex{}

func GetCoordinates(cities []string) (*[]model.ResponseCoordinate, error) {
	//params := c.Request.URL.Query()
	//	fmt.Println("params is", params)
	var coordinates []model.ResponseCoordinate
	for _, city := range cities {

		//fmt.Println("city is", city)
		var singleCord model.Coordinate

		apiUrl := fmt.Sprintf("https://trueway-geocoding.p.rapidapi.com/Geocode?address=%s", city)

		err := utils.RapidApiCall(apiUrl, &singleCord, "trueway-geocoding.p.rapidapi.com")
		if err != nil {
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
	apiUrl := fmt.Sprintf("https://meteostat.p.rapidapi.com/point/%s?lat=%v&lon=%v&start=%v&end=%v&alt=350", Period, Latitude, Longitude, startDate, endDate)

	//fmt.Println("apiUrl", apiUrl)

	err := utils.RapidApiCall(apiUrl, &data, "meteostat.p.rapidapi.com")
	if err != nil {
		return nil, err
	}

	return &data, nil

}
func Daily(ctx *gin.Context, weatherRequest request.WeatherRequest) {

	start := time.Now()

	//get data from params
	cities := weatherRequest.City
	//fmt.Println("params is:", params)
	var totalData []model.Data
	//get start and end date from params'

	startDate := weatherRequest.StartDate
	endDate := weatherRequest.EndDate
	//first get the coordinates
	coordinates, err := GetCoordinates(cities)

	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	//fmt.Println("coordinates are:", coordinates)
	for _, coordinate := range *coordinates {
		wg.Add(1)
		go func(lat, lon float64, city string) {
			data, err := ExtractingData(lat, lon, startDate, endDate, "daily")
			if err != nil {
				response.ErrorResponse(ctx, 400, err.Error())
				wg.Done()
				return
			}
			tempData := model.Data{
				Cityname: city,
				Info:     *data,
			}
			totalData = append(totalData, tempData)
			wg.Done()
		}(coordinate.Latitude, coordinate.Longitude, coordinate.Name)
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
	response.ShowResponse("Success", 200, "Daily fetched successfully", totalData, ctx)
}
func Weekly(ctx *gin.Context, weatherRequest request.WeatherRequest) {

	cities := weatherRequest.City
	//fmt.Println("params is:", params)
	//get start and end date from params'
	var DataSlice []model.GroupData
	startDate := weatherRequest.StartDate
	endDate := weatherRequest.EndDate
	//first get the coordinates
	coordinates, err := GetCoordinates(cities)

	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	for _, coordinate := range *coordinates {
		wg.Add(1)
		go func(lat, lon float64, city string) {
			defer wg.Done()

			data, err := ExtractingData(lat, lon, startDate, endDate, "daily")
			if err != nil {
				response.ErrorResponse(ctx, 400, err.Error())
				return
			}

			var tempSLice []model.TempData
			counter := 1
			//fmt.Println("length of data is", len(data.Resp)/7)

			for i := 0; i < len(data.Resp); i += 7 {
				tempAvg := 0.0
				tempAdd := 0.0

				for j := i; j < i+7 && j < len(data.Resp); j++ {
					tempAdd += data.Resp[j].TempAvg
				}
				tempAvg = tempAdd / 7
				temp := model.TempData{
					Name:    "Week" + strconv.Itoa(counter),
					TempAvg: utils.RoundFloat(tempAvg, 1),
				}
				tempSLice = append(tempSLice, temp)
				counter++
			}

			monthly := model.GroupData{
				CityName: city,
			}
			monthly.Info.Data = tempSLice
			mut.Lock()
			DataSlice = append(DataSlice, monthly)
			mut.Unlock()
		}(coordinate.Latitude, coordinate.Longitude, coordinate.Name)
	}
	wg.Wait()
	response.ShowResponse("Success", 200, "Weekly fetched successfully", DataSlice, ctx)
}

func Monthly(ctx *gin.Context, weatherRequest request.WeatherRequest) {
	cities := weatherRequest.City
	//fmt.Println("params is:", params)
	var totalData []model.Data
	//get start and end date from params'

	startDate := weatherRequest.StartDate
	endDate := weatherRequest.EndDate
	//first get the coordinates
	coordinates, err := GetCoordinates(cities)
	//fmt.Println("coordinates is", coordinates)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	for _, coordinate := range *coordinates {
		wg.Add(1)
		go func(lat, lon float64, city string) {
			data, err := ExtractingData(lat, lon, startDate, endDate, "monthly")
			if err != nil {
				response.ErrorResponse(ctx, 400, err.Error())
				wg.Done()
				return
			}
			tempData := model.Data{
				Cityname: city,
				Info:     *data,
			}
			totalData = append(totalData, tempData)
			wg.Done()
		}(coordinate.Latitude, coordinate.Longitude, coordinate.Name)
	}
	wg.Wait()
	response.ShowResponse("Success", 200, "Monthly fetched successfully", totalData, ctx)
}

func Yearly(ctx *gin.Context, weatherRequest request.WeatherRequest) {
	//get data from params
	cities := weatherRequest.City
	//fmt.Println("params is:", params)
	//get start and end date from params'

	var DataSlice []model.GroupData
	startDate := weatherRequest.StartDate
	endDate := weatherRequest.EndDate
	//first get the coordinates
	coordinates, err := GetCoordinates(cities)

	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	for _, coordinate := range *coordinates {
		wg.Add(1)
		go func(lat, lon float64, city string) {
			data, err := ExtractingData(lat, lon, startDate, endDate, "monthly")
			if err != nil {
				response.ErrorResponse(ctx, 400, err.Error())
				return
			}
			var tempSLice []model.TempData
			counter := 1
			//fmt.Println("length of data is", len(Data.Resp))
			for i := 0; i < len(data.Resp); i += 12 {
				tempAvg := 0.0
				tempAdd := 0.0
				for j := i; j < i+12 && j < len(data.Resp); j++ {
					tempAdd += data.Resp[j].TempAvg
				}
				tempAvg = tempAdd / 12

				temp := model.TempData{
					Name:    "Year " + strconv.Itoa(counter),
					TempAvg: utils.RoundFloat(tempAvg, 1),
				}
				tempSLice = append(tempSLice, temp)
				counter++

			}
			yearly := model.GroupData{
				CityName: city,
			}
			yearly.Info.Data = tempSLice
			mut.Lock()
			DataSlice = append(DataSlice, yearly)
			mut.Unlock()
		}(coordinate.Latitude, coordinate.Longitude, coordinate.Name)
	}
	wg.Wait()
	response.ShowResponse("Success", 200, "Yearly fetched successfully", DataSlice, ctx)
}
