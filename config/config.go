package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	PostgresDSN string
}

// LoadConfig loads the application configuration from the .env file.
func LoadConfig() *AppConfig {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &AppConfig{
		PostgresDSN: "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") +
			" password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") +
			" sslmode=disable",
	}
}
