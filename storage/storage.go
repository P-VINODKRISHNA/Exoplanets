package storage

import (
	"exoplanet-service/models"
	"sync"
)

var (
	exoplanets = make(map[int]models.Exoplanet)
	idCounter  = 1
	mu         sync.Mutex
)

func AddExoplanet(e models.Exoplanet) models.Exoplanet {
	mu.Lock()
	defer mu.Unlock()
	e.ID = idCounter
	exoplanets[idCounter] = e
	idCounter++
	return e
}

func GetAllExoplanets() []models.Exoplanet {
	mu.Lock()
	defer mu.Unlock()
	exoplanetList := make([]models.Exoplanet, 0, len(exoplanets))
	for _, e := range exoplanets {
		exoplanetList = append(exoplanetList, e)
	}
	return exoplanetList
}

func GetExoplanetByID(id int) (models.Exoplanet, bool) {
	mu.Lock()
	defer mu.Unlock()
	e, exists := exoplanets[id]
	return e, exists
}

func UpdateExoplanet(id int, e models.Exoplanet) (models.Exoplanet, bool) {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := exoplanets[id]; exists {
		e.ID = id
		exoplanets[id] = e
		return e, true
	}
	return models.Exoplanet{}, false
}

func DeleteExoplanet(id int) bool {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := exoplanets[id]; exists {
		delete(exoplanets, id)
		return true
	}
	return false
}
