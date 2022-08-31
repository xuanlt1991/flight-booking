package requests

type GetCustomerRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

type ListCustomerRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type CreatedCustomerRequest struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	License     string `json:"license" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
}
