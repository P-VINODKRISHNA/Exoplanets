package model

import (
	"Exoplanet/config"
	"Exoplanet/variables"
	"errors"
	"fmt"
	"log"
	"strings"
)

func UpdateExoPlanet(updates variables.UpdateExoPlanetRequest) (err error) {
	var setClauses []string
	var args []interface{}
	argIdx := 1

	if updates.NewName != "" {
		setClauses = append(setClauses, fmt.Sprintf("name = $%d", argIdx))
		args = append(args, updates.NewName)
		argIdx++
	}
	if updates.Description != "" {
		setClauses = append(setClauses, fmt.Sprintf("description = $%d", argIdx))
		args = append(args, updates.Description)
		argIdx++
	}
	if updates.DistanceFromEarth != 0.0 {
		setClauses = append(setClauses, fmt.Sprintf("distance = $%d", argIdx))
		args = append(args, updates.DistanceFromEarth)
		argIdx++
	}
	if updates.Radius != 0.0 {
		setClauses = append(setClauses, fmt.Sprintf("radius = $%d", argIdx))
		args = append(args, updates.Radius)
		argIdx++
	}
	if updates.Mass != 0.0 {
		setClauses = append(setClauses, fmt.Sprintf("mass = $%d", argIdx))
		args = append(args, updates.Mass)
		argIdx++
	}
	if updates.ExoPlanetType != "" {
		setClauses = append(setClauses, fmt.Sprintf("type = $%d", argIdx))
		args = append(args, updates.ExoPlanetType)
		argIdx++
	}

	if len(setClauses) == 0 {
		return errors.New("no fields to update")
	}

	query := fmt.Sprintf(`
            UPDATE exoPlanets
            SET %s
            WHERE name = $%d
        `, strings.Join(setClauses, ", "), argIdx)

	args = append(args, updates.Name)

	result, err := config.DB.Exec(query, args...)
	if err != nil {
		log.Println(err.Error())
		return errors.New("internal server error")
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		return errors.New("error retrieving affected rows")
	}

	if affectedRows == 0 {
		log.Println("No rows were updated.")
		return errors.New("no rows were updated")
	}

	log.Println("Exoplanet updated successfully.")
	return
}
