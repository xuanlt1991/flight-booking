package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/xuanlt1991/flight-booking/grpc/customer/requests"
	"github.com/xuanlt1991/flight-booking/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *CustomerHandler) ViewCustomer(ctx context.Context, in *pb.ViewCustomerRequest) (*pb.CustomerResponse, error) {
	id, err := uuid.Parse(in.Id)

	if err != nil {
		return nil, err
	}

	req := &requests.GetCustomerRequest{
		Id: id,
	}
	customer, err := h.customerRepository.ViewCustomer(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return customer.ToGRPResponse(), nil

}
