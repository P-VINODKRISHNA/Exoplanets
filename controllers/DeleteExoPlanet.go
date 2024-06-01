package controllers

import (
	"Exoplanet/model"
	"Exoplanet/variables"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteExoPlanetController(c *gin.Context) {
	var err error
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
	err = model.DeleteExoPlanet(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the list of exoplanets
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Exoplanets deleted successfully",
	})

}
