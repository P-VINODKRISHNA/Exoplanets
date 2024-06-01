package model

import (
	"Exoplanet/config"
	"Exoplanet/variables"
	"database/sql"
	"errors"
	"log"
)

func AddExoPlanet(data variables.AddExoPlanetRequest) (err error) {
	var result sql.Result

	result, err = config.DB.Exec(`
            INSERT INTO exoPlanets (name, description, distance, radius, mass, type,status) 
            VALUES ($1, $2, $3, $4, $5, $6,$7)
        `, data.Name, data.Description, data.DistanceFromEarth, data.Radius, data.Mass, data.ExoPlanetType, 1)

	if err != nil {
		log.Println(err.Error())
		err = errors.New("internal server error")
		return
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		err = errors.New("error retrieving affected rows")
		return
	}

	if affectedRows == 0 {
		log.Println("No rows were updated.")
		err = errors.New("no rows were updated")
		return
	}

	log.Println("Exoplanet added successfully.")
	return
}
