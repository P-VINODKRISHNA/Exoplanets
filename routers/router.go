package routers

import (
	"Exoplanet/controllers"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Init() *gin.Engine {

	if strings.EqualFold(viper.GetString("RUN_MODE"), "PROD") {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	v1 := r.Group(viper.GetString("MAIN_PATH"))
	v1.POST(viper.GetString("ADD_EXOPlANET_PATH"), controllers.AddExoPlanetControler)
	v1.GET(viper.GetString("VIEW_EXOPlANET_PATH"), controllers.ListAllExoplanets)
	v1.GET(viper.GetString("VIEW_EXOPlANET_BY_NAME_PATH"), controllers.ListExoPlanetByNameControler)
	v1.PUT(viper.GetString("UPDATE_EXOPlANET_BY_NAME_PATH"), controllers.UpdateExoPlanet)
	v1.PUT(viper.GetString("DELETE_EXOPlANET_BY_NAME_PATH"), controllers.DeleteExoPlanetController)
	v1.POST(viper.GetString("CALCULATE_FUEL_ESTIMATION_PATH"), controllers.FuelEstimationControler)

	return r
}
