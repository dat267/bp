package config

import (
	"encoding/json"
	"os"
	"testing"
)

func TestLoadConfig_EnvPrecedence(t *testing.T) {
	os.Setenv("BP_PORT", "9999")
	defer os.Unsetenv("BP_PORT")

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

func TestLoadConfig_CorruptJSON(t *testing.T) {
	configPath := "corrupt.json"
	os.WriteFile(configPath, []byte("{ invalid json }"), 0644)
	defer os.Remove(configPath)

	// Should not panic, should fallback to defaults
	cfg := LoadConfig(configPath)
	if cfg == nil {
		t.Fatal("Expected config to be initialized even with corrupt file")
	}
	if cfg.Port != "8080" {
		t.Errorf("Expected default Port 8080, got %s", cfg.Port)
	}
}

func TestLoadConfig_AutoGenerate(t *testing.T) {
	// The auto-generation logic only triggers if path is "" (which sets config.json)
	configPath := "config.json"
	os.Remove(configPath)
	defer os.Remove(configPath)

	cfg := LoadConfig("")
	if cfg == nil {
		t.Fatal("Expected config to be initialized")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Error("Expected config file to be auto-generated")
	}
}

func TestConfig_Save(t *testing.T) {
	configPath := "test_save.json"
	defer os.Remove(configPath)

	cfg := &Config{Port: "1234", AppEnv: "test"}
	err := cfg.Save(configPath)
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	// Verify content
	data, _ := os.ReadFile(configPath)
	var loaded Config
	json.Unmarshal(data, &loaded)
	if loaded.Port != "1234" {
		t.Errorf("Expected Port 1234, got %s", loaded.Port)
	}
}

func TestConfig_Validators(t *testing.T) {
	if err := ValidatePort("8080"); err != nil {
		t.Errorf("Valid port failed: %v", err)
	}
	if err := ValidatePort("abc"); err == nil {
		t.Error("Invalid port (string) should fail")
	}
	if err := ValidatePort("70000"); err == nil {
		t.Error("Port out of range should fail")
	}

	if err := ValidateNotEmpty("something"); err != nil {
		t.Errorf("Valid string failed: %v", err)
	}
	if err := ValidateNotEmpty(""); err == nil {
		t.Error("Empty string should fail")
	}
}

func TestConfig_GetMetadata(t *testing.T) {
	cfg := &Config{}
	meta := cfg.GetMetadata()

	if len(meta) != 3 {
		t.Errorf("Expected 3 fields, got %d", len(meta))
	}

	found := false
	for _, f := range meta {
		if f.Name == "Port" && f.Label == "Port" && f.Validator == "port" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected metadata for Port not found or incorrect")
	}
}
