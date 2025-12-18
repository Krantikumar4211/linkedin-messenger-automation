package auth

import (
	"log"
	"time"
)

func Login(email string) error {
	log.Println("Attempting login for:", email)

	// Simulate login delay
	time.Sleep(2 * time.Second)

	log.Println("Login successful")
	return nil
}
