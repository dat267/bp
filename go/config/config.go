package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	AppEnv string `json:"app_env"`
	Port   string `json:"port"`
	APIKey string `json:"api_key"`
}

func LoadConfig(configPath string) *Config {
	cfg := &Config{
		AppEnv: "development",
		Port:   "8080",
		APIKey: "",
	}

	// 1. Load from file (lowest priority)
	if configPath == "" {
		configPath = "config.json"
	}

	if data, err := os.ReadFile(configPath); err == nil {
		json.Unmarshal(data, cfg)
	}

	// 2. Load from environment variables (overrides file/defaults)
	if val := getEnv("APP_ENV"); val != "" {
		cfg.AppEnv = val
	}
	if val := getEnv("PORT"); val != "" {
		cfg.Port = val
	}
	if val := getEnv("API_KEY"); val != "" {
		cfg.APIKey = val
	}

	return cfg
}

func getEnv(key string) string {
	return os.Getenv(key)
}
