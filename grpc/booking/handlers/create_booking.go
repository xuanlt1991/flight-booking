package handlers

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/xuanlt1991/flight-booking/grpc/booking/requests"
	"github.com/xuanlt1991/flight-booking/pb"
)

func (b *BookingHandler) CreateBooking(ctx context.Context, in *pb.BookingRequest) (*pb.BookingResponse, error) {
	log.Printf("Booking Request Before: %v\n", in)

	flight, err := getFlightInfomation(in.FlightId, b.config)

	log.Printf("Booking Request After: %v\n", in)
	if err != nil {
		return nil, err
	}

	bookingDate := time.Now()
	flightDate := time.Date(int(flight.Flight.DepartureDate.Year),
		time.Month(flight.Flight.DepartureDate.Month),
		int(flight.Flight.DepartureDate.Day), flight.Flight.DepartureTime.AsTime().Hour(),
		flight.Flight.DepartureTime.AsTime().Minute(), 0, 0, time.UTC)

	if bookingDate.Sub(flightDate).Hours() >= 12 {
		return nil, errors.New("MUST BE BOOKING 12 HOURS EARLY")
	}

	if flight.Flight.AvailableSlot <= 0 {
		return nil, errors.New("NO AVAILABLE SLOT")
	}

	_, err = getCustomerInfomation(in.CustomerId, b.config)

	if err != nil {
		return nil, err
	}

	cid, err := uuid.Parse(in.CustomerId)

	if err != nil {
		return nil, err
	}

	fid, err := uuid.Parse(in.FlightId)

	if err != nil {
		return nil, err
	}

	booking, err := b.bookingRepository.CreateBooking(ctx, &requests.BookingRequest{
		CustomerId: cid,
		FlightId:   fid,
	})

	if err != nil {
		return nil, err
	}

	flight.Flight.AvailableSlot = flight.Flight.AvailableSlot - 1
	err = updateFlightAvailableSlot(flight.Flight, b.config)
	if err != nil {
		return nil, err
	}

	return booking.ToGRPCResponse(), nil
}
