package handlers

import (
	"context"

	"github.com/xuanlt1991/flight-booking/grpc/customer/requests"
	"github.com/xuanlt1991/flight-booking/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *CustomerHandler) CreateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	req := &requests.CustomerRequest{
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		Address:     in.Address,
		License:     in.License,
		PhoneNumber: in.PhoneNumber,
		Email:       in.Email,
		Password:    in.Password,
	}
	customer, err := h.customerRepository.CreateCustomer(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return customer.ToGRPResponse(), nil
}
