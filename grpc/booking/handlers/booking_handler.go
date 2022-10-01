package handlers

import (
	"context"
	"log"
	"time"

	"github.com/xuanlt1991/flight-booking/config"
	"github.com/xuanlt1991/flight-booking/grpc/booking/repositories"
	"github.com/xuanlt1991/flight-booking/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BookingHandler struct {
	pb.UnimplementedBookingServiceServer
	bookingRepository repositories.BookingRepository
	config            *config.Config
}

func NewBookingHandler(bookingRepository repositories.BookingRepository, config *config.Config) (*BookingHandler, error) {
	return &BookingHandler{bookingRepository: bookingRepository, config: config}, nil
}

func getCustomerInfomation(id string, conf *config.Config) (*pb.CustomerResponse, error) {
	log.Printf("Customer Id: %v\n", id)
	grpcAddress := conf.GRPCConfig.CustomerGRPCServer.Host + ":" + conf.GRPCConfig.CustomerGRPCServer.Port
	customerConn, err := grpc.Dial(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	customerClient := pb.NewCustomerServiceClient(customerConn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	customer, err := customerClient.ViewCustomer(ctx, &pb.ViewCustomerRequest{Id: id})
	if err != nil {
		return nil, err
	}
	log.Printf("customer: %v\n", customer)

	return customer, nil
}

func getFlightInfomation(id string, conf *config.Config) (*pb.FlightResponse, error) {
	log.Printf("Flight Id: %v\n", id)

	grpcAddress := conf.GRPCConfig.FlightGRPCServer.Host + ":" + conf.GRPCConfig.FlightGRPCServer.Port
	flightConn, err := grpc.Dial(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	flightClient := pb.NewFlightServiceClient(flightConn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flight, err := flightClient.ViewFlight(ctx, &pb.ViewFlightRequest{Id: id})
	if err != nil {
		return nil, err
	}
	log.Printf("Flight: %v\n", flight)

	return flight, nil
}

func updateFlightAvailableSlot(flight *pb.Flight, conf *config.Config) error {
	grpcAddress := conf.GRPCConfig.FlightGRPCServer.Host + ":" + conf.GRPCConfig.FlightGRPCServer.Port
	flightConn, err := grpc.Dial(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	filghtClient := pb.NewFlightServiceClient(flightConn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.FlightRequest{
		Id:            flight.Id,
		Name:          flight.Name,
		From:          flight.From,
		To:            flight.To,
		AvailableSlot: flight.AvailableSlot,
		DepartureDate: flight.DepartureDate,
		ArrivalDate:   flight.ArrivalDate,
		DepartureTime: flight.DepartureTime,
		ArrivalTime:   flight.ArrivalTime,
	}
	_, err = filghtClient.UpdateFlight(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
