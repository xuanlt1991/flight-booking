package models

import "time"

type Booking struct {
	Id          int64     `json:"id"`
	CustomerId  int64     `json:"customer_id"`
	FlightId    int64     `json:"flight_id"`
	BookingCode string    `json:"booking_code"`
	Status      string    `json:"status"`
	BookedDate  time.Time `json:"booked_date"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}
