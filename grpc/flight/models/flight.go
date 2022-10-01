package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/xuanlt1991/flight-booking/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Flight struct {
	Id            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name          string    `gorm:"column:name"`
	From          string    `gorm:"column:flight_from"`
	To            string    `gorm:"column:flight_to"`
	Status        string    `gorm:"column:status"`
	AvailableSlot int64     `gorm:"column:available_slot"`
	DepatureDate  time.Time `gorm:"column:departure_date"`
	ArrivalDate   time.Time `gorm:"column:arrival_date"`
	DepartureTime time.Time `gorm:"column:departure_time"`
	ArrivalTime   time.Time `gorm:"column:arrival_time"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	ModifiedAt    time.Time `gorm:"column:modified_at"`
}

func (m *Flight) ToGRPResponse() *pb.FlightResponse {
	return &pb.FlightResponse{
		Flight: &pb.Flight{
			Id:            m.Id.String(),
			Name:          m.Name,
			From:          m.From,
			To:            m.To,
			Status:        m.Status,
			AvailableSlot: m.AvailableSlot,
			DepartureDate: &pb.Date{
				Year:  int32(m.DepatureDate.UTC().Year()),
				Month: int32(m.DepatureDate.UTC().Month()),
				Day:   int32(m.DepatureDate.UTC().Day()),
			},
			ArrivalDate: &pb.Date{
				Year:  int32(m.ArrivalDate.UTC().Year()),
				Month: int32(m.ArrivalDate.UTC().Month()),
				Day:   int32(m.ArrivalDate.UTC().Day()),
			},
			DepartureTime: timestamppb.New(m.DepartureTime),
			ArrivalTime:   timestamppb.New(m.ArrivalTime),
			Audit: &pb.Audit{
				ModifiedAt: timestamppb.New(m.ModifiedAt),
				CreatedAt:  timestamppb.New(m.CreatedAt),
			},
		},
	}
}
