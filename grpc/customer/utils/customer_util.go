package utils

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(rawpassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawpassword), bcrypt.DefaultCost)
	log.Printf("rawpassword: %v - hashedPassword: %v\n", rawpassword, hashedPassword)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func MatchPassword(rawpassword, hashPwd string) bool {
	bytes, err := HashPassword(rawpassword)
	log.Printf("old password: %v\n - hash password: %v\n", rawpassword, bytes)
	log.Printf("stored password: %v\n", hashPwd)
	if err == nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(bytes))

	return err == nil
}

// CheckPassword checks if the provided password is correct or not
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
