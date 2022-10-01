package response

import (
	"time"

	"github.com/google/uuid"
)

type FlightResponse struct {
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
