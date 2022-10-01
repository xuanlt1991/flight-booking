package responses

import "time"

type CustomerResponse struct {
	Id          string    `json:"id"`
	FirstName   string    `json:"first_name" binding:"required"`
	LastName    string    `json:"last_name" binding:"required"`
	Address     string    `json:"address" binding:"required"`
	License     string    `json:"license" binding:"required"`
	PhoneNumber string    `json:"phone_number" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
	Status      string    `json:"status" binding:"status"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}
