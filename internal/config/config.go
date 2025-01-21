package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	AccessTokenSecret  string
	RefreshTokenSecret string
	DatabaseURL        string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	return &Config{
		Port:               GetEnvOrDefault("PORT", "8080"),
		AccessTokenSecret:  GetEnvOrDefault("ACCESS_TOKEN_SECRET", "secret"),
		RefreshTokenSecret: GetEnvOrDefault("REFRESH_TOKEN_SECRET", "secret"),
		DatabaseURL:        GetEnvOrDefault("DATABASE_URL", ""),
	}, nil
}

func GetEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
