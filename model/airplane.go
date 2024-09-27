package model

import "time"

// Airplane es la estructura del modelo de avión.
type Airplane struct {
	ID             string
	Airline        string
	PassengerCount int
	CreationDate   time.Time
	UpdateDate     time.Time
}
