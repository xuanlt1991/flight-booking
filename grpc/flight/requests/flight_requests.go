package requests

import (
	"time"

	"github.com/google/uuid"
)

type FlightRequest struct {
	Id            uuid.UUID
	Name          string
	From          string
	To            string
	Status        string
	AvailableSlot int64
	DepatureDate  time.Time
	ArrivalDate   time.Time
	DepartureTime time.Time
	ArrivalTime   time.Time
}

type SearchFlightRequest struct {
	From         string
	To           string
	DepatureDate time.Time
	ArrivalDate  time.Time
}

type GetFlightRequest struct {
	Id uuid.UUID
}
