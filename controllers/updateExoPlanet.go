package controllers

import (
	"Exoplanet/model"
	"Exoplanet/variables"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateExoPlanet(c *gin.Context) {
	var err error

	// Validate API call

	// Bind JSON request body to checkuserExistRequest struct
	var req variables.UpdateExoPlanetRequest
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

	err = model.UpdateExoPlanet(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "Exoplanet updated successfully"})

	}

	return
}
