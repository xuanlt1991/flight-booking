package handlers

import (
	"github.com/xuanlt1991/flight-booking/grpc/flight/repositories"
	"github.com/xuanlt1991/flight-booking/pb"
)

type FlightHandler struct {
	pb.UnimplementedFlightServiceServer
	flightRepository repositories.FlightRepository
}

func NewFlightHandler(flightRepository repositories.FlightRepository) (*FlightHandler, error) {
	return &FlightHandler{flightRepository: flightRepository}, nil
}
