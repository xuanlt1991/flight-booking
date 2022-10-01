package requests

import "github.com/google/uuid"

type GetCustomerRequest struct {
	Id uuid.UUID
}

// type ListCustomerRequest struct {
// 	PageID   int32 `form:"page_id" binding:"required,min=1"`
// 	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
// }

type CustomerRequest struct {
	Id          uuid.UUID
	FirstName   string
	LastName    string
	Address     string
	License     string
	PhoneNumber string
	Email       string
	Password    string
}

type ChangePasswordRequest struct {
	Id          uuid.UUID
	OldPassword string
	NewPassword string
}
