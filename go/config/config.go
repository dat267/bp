package config

import (
	"encoding/json"
	"os"
	"strings"
)

type Config struct {
	AppEnv string `json:"app_env" flag:"app-env"`
	Port   string `json:"port" flag:"port"`
	APIKey string `json:"api_key" flag:"api-key"`
}

const EnvPrefix = "BP_"

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
	// rclone style: auto-mapping based on prefix + flag name
	if val := getEnvAuto("app-env"); val != "" {
		cfg.AppEnv = val
	}
	if val := getEnvAuto("port"); val != "" {
		cfg.Port = val
	}
	if val := getEnvAuto("api-key"); val != "" {
		cfg.APIKey = val
	}

	return cfg
}

func (cfg *Config) Save(configPath string) error {
	if configPath == "" {
		configPath = "config.json"
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0644)
}

func getEnvAuto(flagName string) string {
	// convert app-env to BP_APP_ENV
	key := EnvPrefix + strings.ToUpper(strings.ReplaceAll(flagName, "-", "_"))
	return os.Getenv(key)
}
