package model

import (
	"Exoplanet/config"
	"Exoplanet/variables"
	"errors"
	"log"
)

func FuelEstimatione(name string) (exoPlanet variables.FuelEstimation, err error) {

	query := `
            SELECT distance, radius, mass, type 
            FROM exoPlanets
            WHERE name = $1
        `
	rows, err := config.DB.Query(query, name)
	if err != nil {
		log.Println(err.Error())
		return exoPlanet, errors.New("internal server error")
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&exoPlanet.DistanceFromEarth, &exoPlanet.Radius, &exoPlanet.Mass, &exoPlanet.ExoPlanetType)
		if err != nil {
			log.Println(err.Error())
			return exoPlanet, errors.New("error scanning row")
		}
	} else {
		return exoPlanet, errors.New("no exoplanet found with the given name")
	}

	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return exoPlanet, errors.New("error iterating over rows")
	}

	log.Println("Exoplanet data retrieved successfully.")
	return exoPlanet, nil
}
