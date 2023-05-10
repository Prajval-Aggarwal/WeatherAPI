package handler

import (
	"main/server/services"

	"github.com/gin-gonic/gin"
)

func DailyHandler(ctx *gin.Context) {
	services.Daily(ctx)
}
func WeeklyHandler(ctx *gin.Context) {
	services.Weekly(ctx)
}
func MonthlyHandler(ctx *gin.Context) {
	services.Monthly(ctx)
}

func YearlyHandler(ctx *gin.Context) {
	services.Yearly(ctx)
}
