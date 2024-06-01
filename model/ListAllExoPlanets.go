package model

import (
	"Exoplanet/config"
	"Exoplanet/variables"
	"errors"
	"log"
)

func ListAllExoPlanets() ([]variables.AddExoPlanetRequest, error) {
	var exoPlanets []variables.AddExoPlanetRequest

	rows, err := config.DB.Query(`
            SELECT name, description, distance, radius, mass, type 
            FROM exoPlanets where status=1
            ORDER BY mass
        `)
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

	log.Println("Exoplanets retrieved and sorted by mass successfully.")
	return exoPlanets, nil
}
