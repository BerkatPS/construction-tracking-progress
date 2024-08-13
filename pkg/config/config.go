package config

import (
	"log"
	"os"
)

type Config struct {
	ServerAddress string
	DatabaseURL   string
	JwtSecret     string
}

func LoadConfig() *Config {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", "localhost:8080"),
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://berkatsaragih:@localhost:5432/construction_track?sslmode=disable"),
		JwtSecret:     getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Using default value for %s: %s", key, defaultValue)
		return defaultValue
	}

	return value
}
