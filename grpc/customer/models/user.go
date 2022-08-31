package models

import "time"

type User struct {
	UserName          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	CustomerId        int64     `json:"customer_id"`
	Status            string    `json:"status"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
	ModifiedAt        time.Time `json:"modified_at"`
}
