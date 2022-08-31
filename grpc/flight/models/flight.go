package models

import "time"

type Flight struct {
	Id            int64     `json:"id"`
	Name          string    `json:"name"`
	From          string    `json:"from"`
	To            string    `json:"to"`
	Status        string    `json:"status"`
	AvailableSlot int64     `json:"available_slot"`
	DepatureDate  time.Time `json:"depature_date"`
	ArrivalDate   time.Time `json:"arrvival_date"`
	DepartureTime string    `json:"departure_time"`
	ArrivalTime   string    `json:"arrival_time"`
	CreatedAt     time.Time `json:"created_at"`
	ModifiedAt    time.Time `json:"modified_at"`
}
