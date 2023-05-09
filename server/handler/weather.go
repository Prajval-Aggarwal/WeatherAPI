package handler

import (
	"main/server/services"

	"github.com/gin-gonic/gin"
)

func DailyHandler(ctx *gin.Context) {
	services.Daily(ctx)
}
func MonthlyHandler(ctx *gin.Context) {

}

func WeeklyHandler(ctx *gin.Context) {

}
func YearlyHandler(ctx *gin.Context) {

}
