package config

import (
	"os"

	"github.com/Thanhbinh1905/seta-training-system/pkg/logger"
	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
	Port        string `mapstructure:"PORT"`
}

func LoadConfig() *Config {
	godotenv.Load()

	postgresHost := os.Getenv("POSTGRES_HOST") // đúng tên biến
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	if postgresHost == "" || postgresPort == "" || postgresUser == "" || postgresPassword == "" || postgresDB == "" {
		logger.Log.Fatal("Missing required environment variables for PostgreSQL")
	}
	databaseUrl := "postgres://" + postgresUser + ":" + postgresPassword + "@" + postgresHost + ":" + postgresPort + "/" + postgresDB + "?sslmode=disable"
	jwt := os.Getenv("JWT_SECRET")
	port := os.Getenv("PORT")

	if jwt == "" || port == "" {
		logger.Log.Fatal("Missing required environment variables")
	}

	return &Config{
		DatabaseURL: databaseUrl,
		JWTSecret:   jwt,
		Port:        port,
	}
}
