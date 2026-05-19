package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("PORT", "9999")
	defer os.Unsetenv("PORT")

	cfg := LoadConfig()
	if cfg.Port != "9999" {
		t.Errorf("Expected Port 9999, got %s", cfg.Port)
	}

	if cfg.AppEnv != "development" {
		t.Errorf("Expected default AppEnv development, got %s", cfg.AppEnv)
	}
}
