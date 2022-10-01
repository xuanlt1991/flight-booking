package responses

import "github.com/google/uuid"

type CustomerResponse struct {
	Id          uuid.UUID
	FirstName   string
	LastName    string
	Address     string
	License     string
	PhoneNumber string
	Email       string
	Status      string
}
