package model

import (
	"Exoplanet/config"
	"Exoplanet/variables"
	"errors"
	"log"
)

func DeleteExoPlanet(data variables.Request) (err error) {
	var query string

	query = `
            UPDATE exoPlanets
            
            SET status =2
            WHERE name = $1 
               
        `

	result, err := config.DB.Exec(query, data.Name)
	if err != nil {
		log.Println(err.Error())
		err = errors.New("internal server error")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		err = errors.New("internal server error")
		return err
	}

	if rowsAffected == 0 {
		return errors.New("No rows updated. Username not found")
	}

	return nil

}
