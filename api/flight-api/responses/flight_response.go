package responses

import (
	"time"
)

type FlightResponse struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	From          string    `json:"from"`
	To            string    `json:"to"`
	Status        string    `json:"status"`
	AvailableSlot int64     `json:"available_slot"`
	DepatureDate  time.Time `json:"depature_date"`
	ArrivalDate   time.Time `json:"arrvival_date"`
	DepartureTime time.Time `json:"departure_time"`
	ArrivalTime   time.Time `json:"arrival_time"`
	CreatedAt     time.Time `json:"created_at"`
	ModifiedAt    time.Time `json:"modified_at"`
}
