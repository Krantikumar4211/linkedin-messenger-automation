package config

import "os"

type Config struct {
	Email     string
	Password  string
	LogLevel  string
	StateFile string
}

func LoadConfig() Config {
	return Config{
		Email:     os.Getenv("LINKEDIN_EMAIL"),
		Password:  os.Getenv("LINKEDIN_PASSWORD"),
		LogLevel:  os.Getenv("LOG_LEVEL"),
		StateFile: os.Getenv("STATE_FILE"),
	}
}
