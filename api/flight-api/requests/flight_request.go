package requests

import (
	"time"
)

type FlightRequest struct {
	Name          string    `json:"name"`
	From          string    `json:"from"`
	To            string    `json:"to"`
	Status        string    `json:"status"`
	AvailableSlot int64     `json:"available_slot"`
	DepatureDate  time.Time `json:"depature_date"`
	ArrivalDate   time.Time `json:"arrival_date"`
	DepartureTime time.Time `json:"departure_time"`
	ArrivalTime   time.Time `json:"arrival_time"`
}

type SearchFlightRequest struct {
	From         string    `json:"from"`
	To           string    `json:"to"`
	DepatureDate time.Time `json:"depature_date"`
	ArrivalDate  time.Time `json:"arrival_date"`
}

// type GetFlightRequest struct {
// 	Id string `json:"id"`
// }
