package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/xuanlt1991/flight-booking/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Booking struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CustomerId  uuid.UUID `gorm:"column:customer_id"`
	FlightId    uuid.UUID `gorm:"column:flight_id"`
	BookingCode string    `gorm:"column:booking_code"`
	Status      string    `gorm:"column:status"`
	BookedDate  time.Time `gorm:"column:booked_date"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	ModifiedAt  time.Time `gorm:"column:modified_at"`
}

func (m *Booking) ToGRPCResponse() *pb.BookingResponse {
	return &pb.BookingResponse{
		Booking: &pb.Booking{
			Id:          m.Id.String(),
			CustomerId:  m.CustomerId.String(),
			FlightId:    m.FlightId.String(),
			BookingCode: m.BookingCode,
			Status:      m.Status,
			BookedDate:  timestamppb.New(m.BookedDate),
			Audit: &pb.Audit{
				ModifiedAt: timestamppb.New(m.ModifiedAt),
				CreatedAt:  timestamppb.New(m.CreatedAt),
			},
		},
	}
}
