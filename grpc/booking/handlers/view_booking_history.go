package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/xuanlt1991/flight-booking/grpc/booking/requests"
	"github.com/xuanlt1991/flight-booking/pb"
)

func (b *BookingHandler) BookingHistory(ctx context.Context, in *pb.ViewBookingHistoryRequest) (*pb.ViewBookingHistoryResponse, error) {
	cid, err := uuid.Parse(in.CustomerId)
	if err != nil {
		return nil, err
	}

	bookings, err := b.bookingRepository.BookingHistory(ctx, &requests.ViewBookingHistoryRequest{CustomerId: cid})

	if err != nil {
		return nil, err
	}
	c, err := getCustomerInfomation(in.CustomerId, b.config)
	if err != nil {
		return nil, err
	}
	res := &pb.ViewBookingHistoryResponse{
		Bookings: []*pb.ViewBookingResponse{},
	}

	for _, booking := range bookings {
		f, err := getFlightInfomation(booking.FlightId.String(), b.config)
		if err != nil {
			return nil, err
		}

		bs := &pb.ViewBookingResponse{
			Booking:  booking.ToGRPCResponse().Booking,
			Customer: c.Customer,
			Flight:   f.Flight,
		}
		res.Bookings = append(res.Bookings, bs)
	}

	return res, nil
}
