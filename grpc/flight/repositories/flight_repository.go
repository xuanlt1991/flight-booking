package repositories

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/xuanlt1991/flight-booking/config"
	"github.com/xuanlt1991/flight-booking/db"
	"github.com/xuanlt1991/flight-booking/grpc/flight/models"
	"github.com/xuanlt1991/flight-booking/grpc/flight/requests"
	"gorm.io/gorm"
)

type flightManager struct {
	*gorm.DB
}

type FlightRepository interface {
	CreateFlight(ctx context.Context, f *requests.FlightRequest) (*models.Flight, error)
	UpdateFlight(ctx context.Context, f *requests.FlightRequest) (*models.Flight, error)
	ViewFlight(ctx context.Context, v *requests.GetFlightRequest) (*models.Flight, error)
	SearchFlight(ctx context.Context, s *requests.SearchFlightRequest) ([]*models.Flight, error)
}

func NewFlightDB(config *config.Config) (FlightRepository, error) {
	db, err := db.OpenDBConnection(config, "flight_booking")

	if err != nil {
		log.Fatal("Cannot connect DB")
	}

	db = db.Debug()

	err = db.AutoMigrate(&models.Flight{})

	if err != nil {
		log.Fatal("Cannot migrate models")
	}

	return &flightManager{db}, nil
}

func (f *flightManager) CreateFlight(ctx context.Context, req *requests.FlightRequest) (*models.Flight, error) {

	flight := &models.Flight{
		Id:            uuid.New(),
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		DepatureDate:  req.DepatureDate,
		ArrivalDate:   req.ArrivalDate,
		Status:        "Active",
		AvailableSlot: req.AvailableSlot,
		DepartureTime: req.DepartureTime,
		ArrivalTime:   req.ArrivalTime,
		CreatedAt:     time.Now(),
		ModifiedAt:    time.Now(),
	}

	if err := f.Create(flight).Error; err != nil {
		log.Println("Cannot create flight")
		return nil, err
	}
	return flight, nil
}
func (f *flightManager) UpdateFlight(ctx context.Context, req *requests.FlightRequest) (*models.Flight, error) {
	flight := &models.Flight{}
	if err := f.First(&flight, req.Id).Error; err != nil {
		log.Fatal("Cannot find customer")
	}
	flight.Name = req.Name
	flight.From = req.From
	flight.To = req.To
	flight.DepatureDate = req.DepatureDate
	flight.ArrivalDate = req.ArrivalDate
	flight.AvailableSlot = req.AvailableSlot
	flight.DepartureTime = req.DepartureTime
	flight.ArrivalTime = req.ArrivalTime

	if err := f.Updates(flight).Error; err != nil {
		return nil, err
	}
	return flight, nil
}

func (f *flightManager) ViewFlight(ctx context.Context, v *requests.GetFlightRequest) (*models.Flight, error) {
	flight := &models.Flight{}

	if err := f.First(&models.Flight{Id: v.Id}).Find(flight).Error; err != nil {
		return nil, err
	}

	return flight, nil
}

func (f *flightManager) SearchFlight(ctx context.Context, s *requests.SearchFlightRequest) ([]*models.Flight, error) {
	flights := []*models.Flight{}

	if err := f.Where("flight_from = ? or flight_to = ? or departure_date = ? or arrival_date = ?",
		s.From, s.To, s.DepatureDate, s.ArrivalDate).Find(&flights).Error; err != nil {
		return nil, err
	}

	return flights, nil
}
