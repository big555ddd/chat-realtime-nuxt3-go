package config

import (
	"app/internal/logger"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	Database()
	app()
	OAuth()
	logger.Init()
	viper.AutomaticEnv()
}

func app() {
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("APP_ENV", "development")

	conf("TOKEN_SECRET_USER", "secret")
	conf("TOKEN_DURATION_USER", 24*time.Hour)
	conf("TOKEN_SECRET_ADMIN", "admin_secret")
	conf("FIRE_BASE_URL", "")
	conf("FIRE_BASE", "")
	conf("PORT", "8080")
}
