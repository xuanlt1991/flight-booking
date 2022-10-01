package requests

type BookingRequest struct {
	Id         string `json:"id"`
	CustomerId string `json:"customer_id"`
	FlightId   string `json:"flight_id"`
}

type ViewBookingRequest struct {
	Id string `json:"id"`
}

type ViewBookingHistoryRequest struct {
	CustomerId string `json:"customer_id"`
}
