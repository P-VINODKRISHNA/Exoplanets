package controllers

import (
	"Exoplanet/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAllExoplanets(c *gin.Context) {
	// Fetch all exoplanets
	exoPlanets, err := model.ListAllExoPlanets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the list of exoplanets
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Exoplanets retrieved successfully",
		"data":    exoPlanets,
	})
}
