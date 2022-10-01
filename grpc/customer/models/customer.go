package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/xuanlt1991/flight-booking/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Customer struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName   string    `gorm:"column:first_name"`
	LastName    string    `gorm:"column:last_name"`
	Address     string    `gorm:"column:address"`
	License     string    `gorm:"column:license"`
	PhoneNumber string    `gorm:"column:phone_number"`
	Email       string    `gorm:"column:email"`
	Status      string    `gorm:"column:status"`
	Password    string    `gorm:"column:password"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	ModifiedAt  time.Time `gorm:"column:modified_at"`
}

func (m *Customer) ToGRPResponse() *pb.CustomerResponse {
	customerRes := &pb.CustomerResponse{
		Customer: &pb.Customer{
			Id:          m.Id.String(),
			FirstName:   m.FirstName,
			LastName:    m.LastName,
			Address:     m.Address,
			License:     m.License,
			PhoneNumber: m.PhoneNumber,
			Email:       m.Email,
			Status:      m.Status,
			Audit: &pb.Audit{
				ModifiedAt: timestamppb.New(m.ModifiedAt),
				CreatedAt:  timestamppb.New(m.CreatedAt),
			},
		},
	}

	log.Printf("customerRes: %v\n", customerRes)
	return customerRes
}
