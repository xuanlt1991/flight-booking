package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/xuanlt1991/flight-booking/grpc/customer/requests"
	"github.com/xuanlt1991/flight-booking/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *CustomerHandler) ChangePassword(ctx context.Context, in *pb.ChangePasswordRequest) (*pb.CustomerResponse, error) {
	id, err := uuid.Parse(in.Id)

	if err != nil {
		return nil, err
	}

	req := &requests.ChangePasswordRequest{
		Id:          id,
		OldPassword: in.OldPassword,
		NewPassword: in.NewPassword,
	}

	customer, err := h.customerRepository.ChangePassword(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return customer.ToGRPResponse(), nil

}
