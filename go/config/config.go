package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	AppEnv string `json:"app_env" flag:"app-env" label:"App Environment"`
	Port   string `json:"port" flag:"port" label:"Port"`
	APIKey string `json:"api_key" flag:"api-key" label:"API Key"`
}

type FieldMeta struct {
	Key       string
	Label     string
	Validator func(string) error
}

func (cfg *Config) GetMetadata() []FieldMeta {
	return []FieldMeta{
		{Key: "AppEnv", Label: "App Environment"},
		{Key: "Port", Label: "Port", Validator: ValidatePort},
		{Key: "APIKey", Label: "API Key", Validator: ValidateNotEmpty},
	}
}

// Helper validators moved to config package for schema reuse
func ValidatePort(val string) error {
	var p int
	_, err := fmt.Sscanf(val, "%d", &p)
	if err != nil {
		return fmt.Errorf("port must be a number")
	}
	if p < 1 || p > 65535 {
		return fmt.Errorf("port must be between 1 and 65535")
	}
	return nil
}

func ValidateNotEmpty(val string) error {
	if strings.TrimSpace(val) == "" {
		return fmt.Errorf("value cannot be empty")
	}
	return nil
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
