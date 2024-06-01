package controllers

import (
	"Exoplanet/model"
	"Exoplanet/variables"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FuelEstimationControler(c *gin.Context) {
	var err error

	// Bind JSON request body to ListExoPlanetByNameRequest struct
	var req variables.Request
	if err = c.ShouldBindJSON(&req); err != nil {
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
	exoPlanets, err := model.FuelEstimatione(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	crewCapacity, err := strconv.Atoi(req.Crew)
	if err != nil || crewCapacity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid crew capacity"})
		return
	}

	var gravity float64
	switch exoPlanets.ExoPlanetType {
	case "GasGiant":
		gravity = 0.5 / (exoPlanets.Radius * exoPlanets.Radius)
	case "Terrestrial":
		gravity = exoPlanets.Mass / (exoPlanets.Radius * exoPlanets.Radius)
	}

	fuel := float64(exoPlanets.DistanceFromEarth) / (gravity * gravity) * float64(crewCapacity)
	// Respond with the list of exoplanets
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": fmt.Sprintf("required fuel %.2f", fuel),
		"Fuel":    fuel,
	})
}
