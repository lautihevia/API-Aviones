package repositories

import (
	"ejemplo1/model"
	"errors"
	"strings"
)

// Simulación de base de datos en memoria
var airplanes = make(map[string]model.Airplane)

// SaveAirplane guarda un avión en la "base de datos".
func SaveAirplane(airplane model.Airplane) error {
	airplanes[airplane.ID] = airplane
	return nil
}

// FindAirplanes busca aviones, con filtrado opcional por aerolínea.
func FindAirplanes(airline string) ([]model.Airplane, error) {
	var result []model.Airplane
	for _, airplane := range airplanes {
		if airline == "" || containsIgnoreCase(airplane.Airline, airline) {
			result = append(result, airplane)
		}
	}
	return result, nil
}

// FindAirplaneByID busca un avión por su ID.
func FindAirplaneByID(id string) (model.Airplane, error) {
	airplane, exists := airplanes[id]
	if !exists {
		return model.Airplane{}, errors.New("airplane not found")
	}
	return airplane, nil
}

// DeleteAirplane elimina un avión por su ID.
func DeleteAirplane(id string) error {
	delete(airplanes, id)
	return nil
}

// containsIgnoreCase verifica si una cadena contiene otra, ignorando mayúsculas y minúsculas.
func containsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}
