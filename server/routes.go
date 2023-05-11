package server

import (
	_ "main/docs"
	"main/server/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {

	router := server.engine.Group("/api")

	//weather routes
	router.POST("/daily", handler.DailyHandler)
	router.POST("/weekly", handler.WeeklyHandler)
	router.POST("/monthly", handler.MonthlyHandler)
	router.POST("/yearly", handler.YearlyHandler)

	//swagger route
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
