package handlers

import (
	"github.com/xuanlt1991/flight-booking/grpc/customer/repositories"
	"github.com/xuanlt1991/flight-booking/pb"
)

type CustomerHandler struct {
	pb.UnimplementedCustomerServiceServer
	customerRepository repositories.CustomerRepository
}

func NewCustomerHandler(customerRepository repositories.CustomerRepository) (*CustomerHandler, error) {
	return &CustomerHandler{customerRepository: customerRepository}, nil
}
