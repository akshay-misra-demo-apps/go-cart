package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var SECRET = "go-cart.com.007"

func Encrypt(key string, cost int) (hashed string, err error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(key), cost)
	return string(hashedBytes), err
}

// IsValidPassword return true if password matched, else false
func IsValidPassword(original, hashed string) bool {
	fmt.Printf("original password: %v - %v", original, hashed)
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(original))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false
	}

	return true
}
