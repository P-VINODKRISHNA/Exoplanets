package main

import (
	"exoplanet-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/Addexoplanets", handlers.AddExoplanetHandler)
	r.GET("/GetALLexoplanets", handlers.ListExoplanetsHandler)
	r.GET("/exoplanetsByID/:id", handlers.GetExoplanetHandler)
	r.PUT("/UpdateexoplanetsByID/:id", handlers.UpdateExoplanetHandler)
	r.DELETE("/DeletexoplanetsByID/:id", handlers.DeleteExoplanetHandler)
	r.GET("/CalculateFuelexoplanets/:id/fuel", handlers.CalculateFuelHandler)

	r.Run(":8080")
}
