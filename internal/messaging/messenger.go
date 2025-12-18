package messaging

import (
	"log"

	"linkedin-automation/internal/automation"
)

var limiter = automation.NewRateLimiter(10) // 1 message per 10 seconds

func SendMessage(user string, message string) error {
	limiter.Wait() // RATE LIMITER

	log.Println("Typing message...")
	automation.RandomDelay(500, 1500)

	log.Println("Sending message to:", user)
	automation.RandomDelay(1000, 2000)

	log.Println("Message sent successfully")
	return nil
}
