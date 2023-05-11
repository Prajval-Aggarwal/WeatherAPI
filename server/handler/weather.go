package handler

import (
	"fmt"
	"main/server/request"
	"main/server/response"
	"main/server/services"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

// @Description	Show the daily historical data from start date to end date
// @Accept			json
// @Produce		json
// @Success		200	{object}	response.Success
// @Failure		400	{object}	response.Error
// @Param weatherDetails body request.WeatherRequest true "Weather Details"
// @Tags			Weather
// @Router			/daily [post]
func DailyHandler(ctx *gin.Context) {
	var weatherRequest request.WeatherRequest

	err := utils.RequestDecoding(ctx, &weatherRequest)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	fmt.Println("cities is:", weatherRequest)
	services.Daily(ctx, weatherRequest)

}

// @Description	Show the weekly historical data from start date to end date
// @Accept			json
// @Produce		json
// @Success		200	{object}	response.Success
// @Failure		400	{object}	response.Error
// @Param weatherDetails body request.WeatherRequest true "Weather De"
// @Tags			Weather
// @Router			/weekly [post]
func WeeklyHandler(ctx *gin.Context) {
	var weatherRequest request.WeatherRequest

	err := utils.RequestDecoding(ctx, &weatherRequest)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	fmt.Println("cities is:", weatherRequest)
	services.Weekly(ctx, weatherRequest)

}

// @Description	Show the monthly historical data from start date to end date
// @Accept			json
// @Produce		json
// @Success		200	{object}	response.Success
// @Failure		400	{object}	response.Error
// @Param weatherDetails body request.WeatherRequest true "Weather De"
// @Tags			Weather
// @Router			/monthly [post]
func MonthlyHandler(ctx *gin.Context) {
	var weatherRequest request.WeatherRequest

	err := utils.RequestDecoding(ctx, &weatherRequest)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	fmt.Println("cities is:", weatherRequest)
	services.Monthly(ctx, weatherRequest)

}

// @Description	Show the yearly historical data from start date to end date
// @Accept			json
// @Produce		json
// @Success		200	{object}	response.Success
// @Failure		400	{object}	response.Error
// @Param weatherDetails body request.WeatherRequest true "Weather De"
// @Tags			Weather
// @Router			/yearly [post]
func YearlyHandler(ctx *gin.Context) {
	var weatherRequest request.WeatherRequest

	err := utils.RequestDecoding(ctx, &weatherRequest)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	fmt.Println("cities is:", weatherRequest)
	services.Yearly(ctx, weatherRequest)

}
