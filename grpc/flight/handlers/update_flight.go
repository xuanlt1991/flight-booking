package handlers

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/xuanlt1991/flight-booking/grpc/flight/requests"
	"github.com/xuanlt1991/flight-booking/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *FlightHandler) UpdateFlight(ctx context.Context, in *pb.FlightRequest) (*pb.FlightResponse, error) {
	id, err := uuid.Parse(in.Id)

	if err != nil {
		log.Fatal(err)
	}

	req := &requests.FlightRequest{
		Id:            id,
		Name:          in.Name,
		From:          in.From,
		To:            in.To,
		AvailableSlot: in.AvailableSlot,
		DepatureDate:  time.Date(int(in.DepartureDate.Year), time.Month(in.DepartureDate.Month), int(in.DepartureDate.Day), 0, 0, 0, 0, time.UTC),
		ArrivalDate:   time.Date(int(in.ArrivalDate.Year), time.Month(in.ArrivalDate.Month), int(in.ArrivalDate.Day), 0, 0, 0, 0, time.UTC),
		DepartureTime: in.DepartureTime.AsTime(),
		ArrivalTime:   in.ArrivalTime.AsTime(),
	}

	flight, err := h.flightRepository.UpdateFlight(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return flight.ToGRPResponse(), nil
}
