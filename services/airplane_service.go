package services

import (
	"ejemplo1/dto"
	"ejemplo1/model"
	"ejemplo1/repositories"
	"ejemplo1/utils"
	"time"
)

// CreateAirplane crea un nuevo avión.
func CreateAirplane(req dto.CreateAirplaneRequest) (dto.AirplaneResponse, error) {
	airplane := model.Airplane{
		ID:             utils.GenerateID(),
		Airline:        req.Airline,
		PassengerCount: req.PassengerCount,
		CreationDate:   time.Now(),
		UpdateDate:     time.Now(),
	}

	if err := repositories.SaveAirplane(airplane); err != nil {
		return dto.AirplaneResponse{}, err
	}

	return dto.AirplaneResponse{
		ID:             airplane.ID,
		Airline:        airplane.Airline,
		PassengerCount: airplane.PassengerCount,
		CreationDate:   airplane.CreationDate,
		UpdateDate:     airplane.UpdateDate,
	}, nil
}

// GetAirplanes obtiene una lista de aviones, con filtrado opcional por aerolínea.
func GetAirplanes(airline string) ([]dto.AirplaneResponse, error) {
	airplanes, err := repositories.FindAirplanes(airline)
	if err != nil {
		return nil, err
	}

	var response []dto.AirplaneResponse
	for _, airplane := range airplanes {
		response = append(response, dto.AirplaneResponse{
			ID:             airplane.ID,
			Airline:        airplane.Airline,
			PassengerCount: airplane.PassengerCount,
			CreationDate:   airplane.CreationDate,
			UpdateDate:     airplane.UpdateDate,
		})
	}

	return response, nil
}

// GetAirplaneByID obtiene un avión por su ID.
func GetAirplaneByID(id string) (dto.AirplaneResponse, error) {
	airplane, err := repositories.FindAirplaneByID(id)
	if err != nil {
		return dto.AirplaneResponse{}, err
	}

	return dto.AirplaneResponse{
		ID:             airplane.ID,
		Airline:        airplane.Airline,
		PassengerCount: airplane.PassengerCount,
		CreationDate:   airplane.CreationDate,
		UpdateDate:     airplane.UpdateDate,
	}, nil
}

// UpdateAirplane actualiza un avión por su ID.
func UpdateAirplane(id string, req dto.UpdateAirplaneRequest) (dto.AirplaneResponse, error) {
	airplane, err := repositories.FindAirplaneByID(id)
	if err != nil {
		return dto.AirplaneResponse{}, err
	}

	airplane.Airline = req.Airline
	airplane.PassengerCount = req.PassengerCount
	airplane.UpdateDate = time.Now()

	if err := repositories.SaveAirplane(airplane); err != nil {
		return dto.AirplaneResponse{}, err
	}

	return dto.AirplaneResponse{
		ID:             airplane.ID,
		Airline:        airplane.Airline,
		PassengerCount: airplane.PassengerCount,
		CreationDate:   airplane.CreationDate,
		UpdateDate:     airplane.UpdateDate,
	}, nil
}

// DeleteAirplane elimina un avión por su ID.
func DeleteAirplane(id string) error {
	return repositories.DeleteAirplane(id)
}
