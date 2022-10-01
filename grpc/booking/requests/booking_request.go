package requests

import (
	"github.com/google/uuid"
)

type BookingRequest struct {
	Id         uuid.UUID
	CustomerId uuid.UUID
	FlightId   uuid.UUID
}

type ViewBookingRequest struct {
	Id uuid.UUID
}

type ViewBookingHistoryRequest struct {
	CustomerId uuid.UUID
}
