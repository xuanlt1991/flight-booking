package responses

import (
	"time"
)

type BookingResponse struct {
	Id          string    `json:"id"`
	CustomerId  string    `json:"customer_id"`
	FlightId    string    `json:"flight_id"`
	BookingCode string    `json:"booking_code"`
	Status      string    `json:"status"`
	BookedDate  time.Time `json:"booked_date"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}
