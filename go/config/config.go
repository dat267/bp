package config

import (
	"os"
)

type Config struct {
	AppEnv  string
	Port    string
	APIKey  string
}

func LoadConfig() *Config {
	return &Config{
		AppEnv: getEnv("APP_ENV", "development"),
		Port:   getEnv("PORT", "8080"),
		APIKey: getEnv("API_KEY", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
