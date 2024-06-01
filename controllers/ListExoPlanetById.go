package controllers

import (
	"Exoplanet/model"
	"Exoplanet/variables"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListExoPlanetByNameControler(c *gin.Context) {
	var err error

	// Bind JSON request body to ListExoPlanetByNameRequest struct
	var req variables.Request
	if err = c.ShouldBindQuery(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	obj, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("Request", string(obj))

	// Retrieve exoplanets by name
	exoPlanets, err := model.ListExoPlanetsByName(req.Name)
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
