package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/xuanlt1991/flight-booking/pb"
)

func (b *BookingHandler) CancelBooking(ctx context.Context, in *pb.CancelBookingRequest) (*pb.Empty, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, err
	}
	err = b.bookingRepository.CancelBooking(ctx, id)

	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
