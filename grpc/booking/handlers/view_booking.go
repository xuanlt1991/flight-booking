package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/xuanlt1991/flight-booking/grpc/booking/requests"
	"github.com/xuanlt1991/flight-booking/pb"
)

func (b *BookingHandler) ViewBooking(ctx context.Context, in *pb.ViewBookingRequest) (*pb.ViewBookingResponse, error) {
	cid, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, err
	}
	booking, err := b.bookingRepository.ViewBooking(ctx, &requests.ViewBookingRequest{Id: cid})
	if err != nil {
		return nil, err
	}
	c, err := getCustomerInfomation(booking.CustomerId.String(), b.config)
	if err != nil {
		return nil, err
	}
	f, err := getFlightInfomation(booking.FlightId.String(), b.config)
	if err != nil {
		return nil, err
	}

	return &pb.ViewBookingResponse{
		Booking:  booking.ToGRPCResponse().Booking,
		Customer: c.Customer,
		Flight:   f.Flight,
	}, nil
}
