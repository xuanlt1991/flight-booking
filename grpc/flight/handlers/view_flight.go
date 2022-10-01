package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/xuanlt1991/flight-booking/grpc/flight/requests"
	"github.com/xuanlt1991/flight-booking/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *FlightHandler) ViewFlight(ctx context.Context, in *pb.ViewFlightRequest) (*pb.FlightResponse, error) {
	id, err := uuid.Parse(in.Id)

	if err != nil {
		return nil, err
	}

	req := &requests.GetFlightRequest{
		Id: id,
	}
	flight, err := h.flightRepository.ViewFlight(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return flight.ToGRPResponse(), nil
}
