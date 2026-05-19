package config

import (
	"os"
	"testing"
)

func TestLoadConfig_EnvPrecedence(t *testing.T) {
	os.Setenv("PORT", "9999")
	defer os.Unsetenv("PORT")

	cfg := LoadConfig("")
	if cfg.Port != "9999" {
		t.Errorf("Expected Port 9999 from env, got %s", cfg.Port)
	}
}

func TestLoadConfig_FileFallback(t *testing.T) {
	content := `{"port": "1234", "app_env": "production"}`
	err := os.WriteFile("config.json", []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("config.json")

	cfg := LoadConfig("")
	if cfg.Port != "1234" {
		t.Errorf("Expected Port 1234 from file, got %s", cfg.Port)
	}
}

func TestLoadConfig_CustomPath(t *testing.T) {
	content := `{"port": "5555"}`
	err := os.WriteFile("custom.json", []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("custom.json")

	cfg := LoadConfig("custom.json")
	if cfg.Port != "5555" {
		t.Errorf("Expected Port 5555 from custom file, got %s", cfg.Port)
	}
}
