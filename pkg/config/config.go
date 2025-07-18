package config

import (
	"os"

	"github.com/Thanhbinh1905/seta-training-system/pkg/logger"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`

	Port string `mapstructure:"PORT"`
}

func LoadConfig() *Config {
	godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	jwt := os.Getenv("JWT_SECRET")
	port := os.Getenv("PORT")

	if dbURL == "" || jwt == "" || port == "" {
		logger.Log.Fatal("Missing required environment variables")
	}

	return &Config{
		DatabaseURL: dbURL,
		JWTSecret:   jwt,
		Port:        port,
	}
}
