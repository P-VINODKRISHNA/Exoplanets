package handlers

import (
	"exoplanet-service/models"
	"exoplanet-service/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddExoplanetHandler(c *gin.Context) {
	var newExoplanet models.Exoplanet
	if err := c.ShouldBindJSON(&newExoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := newExoplanet.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addedExoplanet := storage.AddExoplanet(newExoplanet)
	c.JSON(http.StatusOK, addedExoplanet)
}

func ListExoplanetsHandler(c *gin.Context) {
	exoplanetList := storage.GetAllExoplanets()
	c.JSON(http.StatusOK, exoplanetList)
}

func GetExoplanetHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	exoplanet, exists := storage.GetExoplanetByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exoplanet not found"})
		return
	}

	c.JSON(http.StatusOK, exoplanet)
}

func UpdateExoplanetHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedExoplanet models.Exoplanet
	if err := c.ShouldBindJSON(&updatedExoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := updatedExoplanet.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exoplanet, exists := storage.UpdateExoplanet(id, updatedExoplanet)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exoplanet not found"})
		return
	}

	c.JSON(http.StatusOK, exoplanet)
}

func DeleteExoplanetHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if !storage.DeleteExoplanet(id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exoplanet not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Exoplanet deleted"})
}

func CalculateFuelHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	exoplanet, exists := storage.GetExoplanetByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exoplanet not found"})
		return
	}

	crewCapacity, err := strconv.Atoi(c.Query("crew"))
	if err != nil || crewCapacity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid crew capacity"})
		return
	}

	var gravity float64
	switch exoplanet.Type {
	case models.GasGiant:
		gravity = 0.5 / (exoplanet.Radius * exoplanet.Radius)
	case models.Terrestrial:
		gravity = exoplanet.Mass / (exoplanet.Radius * exoplanet.Radius)
	}

	fuel := float64(exoplanet.Distance) / (gravity * gravity) * float64(crewCapacity)
	c.JSON(http.StatusOK, gin.H{"fuel": fuel})
}
