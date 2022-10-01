package repositories

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/xuanlt1991/flight-booking/config"
	"github.com/xuanlt1991/flight-booking/db"
	"github.com/xuanlt1991/flight-booking/grpc/booking/models"
	"github.com/xuanlt1991/flight-booking/grpc/booking/requests"
	"github.com/xuanlt1991/flight-booking/util"
	"gorm.io/gorm"
)

type bookingManager struct {
	*gorm.DB
}

type BookingRepository interface {
	CreateBooking(ctx context.Context, in *requests.BookingRequest) (*models.Booking, error)
	ViewBooking(ctx context.Context, in *requests.ViewBookingRequest) (*models.Booking, error)
	BookingHistory(ctx context.Context, in *requests.ViewBookingHistoryRequest) ([]*models.Booking, error)
	CancelBooking(ctx context.Context, id uuid.UUID) error
}

func NewBookingDB(config *config.Config) (BookingRepository, error) {
	db, err := db.OpenDBConnection(config, "flight_booking")

	if err != nil {
		log.Fatal("Cannot connect DB")
	}

	db = db.Debug()

	err = db.AutoMigrate(&models.Booking{})

	if err != nil {
		log.Fatal("Cannot migrate models")
	}

	return &bookingManager{db}, nil
}

func (b *bookingManager) CreateBooking(ctx context.Context, req *requests.BookingRequest) (*models.Booking, error) {
	booking := &models.Booking{
		Id:          uuid.New(),
		CustomerId:  req.CustomerId,
		FlightId:    req.FlightId,
		BookingCode: util.RandomString(10),
		Status:      "Active",
		BookedDate:  time.Now(),
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}

	if err := b.Create(booking).Error; err != nil {
		log.Println("Cannot create flight booking")
		return nil, err
	}
	return booking, nil
}

func (b *bookingManager) ViewBooking(ctx context.Context, req *requests.ViewBookingRequest) (*models.Booking, error) {
	booking := &models.Booking{}

	if err := b.First(&models.Booking{Id: req.Id}).Find(&booking).Error; err != nil {
		return nil, err
	}

	return booking, nil
}
func (b *bookingManager) BookingHistory(ctx context.Context, req *requests.ViewBookingHistoryRequest) ([]*models.Booking, error) {
	log.Printf("BookingHistory was invoked with id: %v\n", req.CustomerId)
	bookings := []*models.Booking{}
	err := b.Where("customer_id = ?", req.CustomerId).Find(&bookings).Error

	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (b *bookingManager) CancelBooking(ctx context.Context, id uuid.UUID) error {
	err := b.Where(&models.Booking{Id: id}).Select("status").Updates(&models.Booking{Id: id, Status: "Canceled"}).Error

	if err != nil {
		return err
	}
	return nil
}
