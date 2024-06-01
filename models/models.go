package models

import "errors"

type ExoplanetType string

const (
	GasGiant    ExoplanetType = "GasGiant"
	Terrestrial ExoplanetType = "Terrestrial"
)

type Exoplanet struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Distance    int           `json:"distance"`       // light years
	Radius      float64       `json:"radius"`         // Earth-radius unit
	Mass        float64       `json:"mass,omitempty"` // Earth-mass unit, only for Terrestrial
	Type        ExoplanetType `json:"type"`
}

func (e *Exoplanet) Validate() error {
	if e.Distance <= 10 || e.Distance >= 1000 {
		return errors.New("distance must be between 10 and 1000 light years")
	}
	if e.Radius <= 0.1 || e.Radius >= 10 {
		return errors.New("radius must be between 0.1 and 10 Earth-radius units")
	}
	if e.Type == Terrestrial && (e.Mass <= 0.1 || e.Mass >= 10) {
		return errors.New("mass must be between 0.1 and 10 Earth-mass units for Terrestrial planets")
	}
	return nil
}
