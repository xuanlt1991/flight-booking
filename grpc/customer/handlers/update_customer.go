package handlers

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/xuanlt1991/flight-booking/grpc/customer/requests"
	"github.com/xuanlt1991/flight-booking/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *CustomerHandler) UpdateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	id, err := uuid.Parse(in.Id)

	if err != nil {
		log.Fatal(err)
	}

	req := &requests.CustomerRequest{
		Id:          id,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		Address:     in.Address,
		License:     in.License,
		PhoneNumber: in.PhoneNumber,
		Email:       in.Email,
		Password:    in.Password,
	}

	customer, err := h.customerRepository.UpdateCustomer(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return customer.ToGRPResponse(), nil
}
