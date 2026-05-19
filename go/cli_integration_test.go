package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestCLIFlagPositions(t *testing.T) {
	// Build the CLI binary first
	build := exec.Command("go", "build", "-o", "cli_test_bin", "./cmd/cli")
	if err := build.Run(); err != nil {
		t.Fatalf("Failed to build CLI: %v", err)
	}
	defer exec.Command("rm", "cli_test_bin").Run()

	tests := []struct {
		name     string
		args     []string
		contains string
	}{
		{"Verbose before", []string{"--verbose", "info"}, "Verbose:     true"},
		{"Verbose after", []string{"info", "--verbose"}, "Verbose:     true"},
		{"Subcommand flag with global", []string{"--verbose", "hello", "--name=Tester"}, "Hello, Tester!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("./cli_test_bin", tt.args...)
			var out bytes.Buffer
			cmd.Stdout = &out
			if err := cmd.Run(); err != nil {
				t.Errorf("Command failed: %v", err)
			}
			if !strings.Contains(out.String(), tt.contains) {
				t.Errorf("Output missing expected string %q: %s", tt.contains, out.String())
			}
		})
	}
}
