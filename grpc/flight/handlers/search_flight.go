package handlers

import (
	"context"
	"time"

	"github.com/xuanlt1991/flight-booking/grpc/flight/requests"
	"github.com/xuanlt1991/flight-booking/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *FlightHandler) SearchFlight(ctx context.Context, in *pb.SearchFlightRequest) (*pb.SearchFlightResponse, error) {
	req := &requests.SearchFlightRequest{
		From:         in.From,
		To:           in.To,
		DepatureDate: time.Date(int(in.DepartureDate.Year), time.Month(in.DepartureDate.Month), int(in.DepartureDate.Day), 0, 0, 0, 0, time.UTC),
		ArrivalDate:  time.Date(int(in.ArrivalDate.Year), time.Month(in.ArrivalDate.Month), int(in.ArrivalDate.Day), 0, 0, 0, 0, time.UTC),
	}

	flights, err := h.flightRepository.SearchFlight(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err != nil {
		return nil, err
	}

	res := &pb.SearchFlightResponse{
		Flights: []*pb.FlightResponse{},
	}

	for _, flight := range flights {
		res.Flights = append(res.Flights, flight.ToGRPResponse())
	}

	return res, nil
}
