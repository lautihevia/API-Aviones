package dto

import "time"

// CreateAirplaneRequest es el DTO para la creación de un avión.
type CreateAirplaneRequest struct {
	Airline        string `json:"airline" binding:"required"`
	PassengerCount int    `json:"passenger_count" binding:"required"`
}

// UpdateAirplaneRequest es el DTO para la actualización de un avión.
type UpdateAirplaneRequest struct {
	Airline        string `json:"airline" binding:"required"`
	PassengerCount int    `json:"passenger_count" binding:"required"`
}

// AirplaneResponse es el DTO para la respuesta de un avión.
type AirplaneResponse struct {
	ID             string    `json:"id"`
	Airline        string    `json:"airline"`
	PassengerCount int       `json:"passenger_count"`
	CreationDate   time.Time `json:"creation_date"`
	UpdateDate     time.Time `json:"update_date"`
}
