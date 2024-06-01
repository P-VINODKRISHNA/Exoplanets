package utils

import (
	"Exoplanet/variables"
	"errors"
)

func ValidateExoPlanet(exoplanet variables.AddExoPlanetRequest) error {
	if exoplanet.DistanceFromEarth <= 10 || exoplanet.DistanceFromEarth >= 1000 {
		return errors.New("distance must be between 10 and 1000 light years")
	}
	if exoplanet.Radius <= 0.1 || exoplanet.Radius >= 10 {
		return errors.New("radius must be between 0.1 and 10 Earth-radius units")
	}
	if exoplanet.ExoPlanetType == "Terrestrial" && (exoplanet.Mass <= 0.1 || exoplanet.Mass >= 10) {
		return errors.New("mass must be between 0.1 and 10 Earth-mass units for Terrestrial planets")
	}
	return nil

}
