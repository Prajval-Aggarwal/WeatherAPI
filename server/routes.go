package server

import (
	_ "main/docs"
	"main/server/handler"
)

func ConfigureRoutes(server *Server) {

	router := server.engine.Group("/api")
	router.GET("/daily", handler.DailyHandler)
	router.GET("/weekly", handler.WeeklyHandler)

	router.GET("/monthly", handler.MonthlyHandler)

	router.GET("/yearly", handler.YearlyHandler)

}
