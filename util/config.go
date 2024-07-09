package util

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DBName string
}

// NewConfig creates a new instance of Config.
//
// If in developement, loades env vars from .env file.
func NewConfig() *Config {
	if os.Getenv("ENV") != "production" {
		_ = godotenv.Load(".env")
	}

	return &Config{
		Port:   getEnvVar("PORT", "3000"),
		DBName: getEnvVar("DB_NAME", "goDB"),
	}
}

// getEnvVar searches for a given key or return the fallback of key doesn't exist.
func getEnvVar(key, fallback string) string {
	if val, exist := os.LookupEnv(key); exist {
		return val
	}
	return fallback
}
