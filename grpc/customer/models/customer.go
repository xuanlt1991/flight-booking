package models

import "time"

type Customer struct {
	Id          int64     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Address     string    `json:"address"`
	License     string    `json:"license"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}
