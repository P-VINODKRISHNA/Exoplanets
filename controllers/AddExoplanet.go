package controllers

import (
	"Exoplanet/model"
	"Exoplanet/utils"
	"Exoplanet/variables"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddExoPlanetControler(c *gin.Context) {
	var err error

	// Validate API call

	// Bind JSON request body to checkuserExistRequest struct
	var req variables.AddExoPlanetRequest
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
	if err = utils.ValidateExoPlanet(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = model.AddExoPlanet(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "Exoplanet added successfully"})

	}

	return
}
