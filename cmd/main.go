package main

import (
	"log"

	"github.com/joho/godotenv"

	"linkedin-automation/internal/auth"
	"linkedin-automation/internal/config"
	"linkedin-automation/internal/logger"
	"linkedin-automation/internal/messaging"
	"linkedin-automation/internal/storage"
)

func main() {
	// Load environment variables from .env file
	_ = godotenv.Load()

	logger.InitLogger()
	cfg := config.LoadConfig()

	log.Println("Starting LinkedIn Automation")

	err := auth.Login(cfg.Email)
	if err != nil {
		log.Fatal("Login failed")
	}

	state := storage.LoadState(cfg.StateFile)

	err = messaging.SendMessage("recruiter_user", "Hello! Thanks for connecting.")
	if err == nil {
		state.SentMessages = append(state.SentMessages, "recruiter_user")
		if err := storage.SaveState(cfg.StateFile, state); err != nil {
			log.Println("Failed to save state:", err)
		}
	}

	log.Println("Automation completed")
}
