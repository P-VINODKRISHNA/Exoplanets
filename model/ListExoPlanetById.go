package model

import (
	"Exoplanet/config"
	"Exoplanet/variables"
	"errors"
	"log"
)

func ListExoPlanetsByName(name string) ([]variables.AddExoPlanetRequest, error) {
	var exoPlanets []variables.AddExoPlanetRequest

	query := `
            SELECT name, description, distance, radius, mass, type 
            FROM exoPlanets
            WHERE name = $1
        `
	rows, err := config.DB.Query(query, name)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("internal server error")
	}
	defer rows.Close()

	for rows.Next() {
		var exoPlanet variables.AddExoPlanetRequest
		err = rows.Scan(&exoPlanet.Name, &exoPlanet.Description, &exoPlanet.DistanceFromEarth, &exoPlanet.Radius, &exoPlanet.Mass, &exoPlanet.ExoPlanetType)
		if err != nil {
			log.Println(err.Error())
			return nil, errors.New("error scanning row")
		}
		exoPlanets = append(exoPlanets, exoPlanet)
	}

	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return nil, errors.New("error iterating over rows")
	}

	log.Println("Exoplanets retrieved successfully.")
	return exoPlanets, nil
}
