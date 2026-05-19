package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"bp/config"
)

type ConfigCommand struct {
	fs     *flag.FlagSet
	config *config.Config
}

func NewConfigCommand(cfg *config.Config) *ConfigCommand {
	return &ConfigCommand{
		fs:     flag.NewFlagSet("config", flag.ContinueOnError),
		config: cfg,
	}
}

func (c *ConfigCommand) Name() string        { return c.fs.Name() }
func (c *ConfigCommand) Description() string { return "Interactively setup configuration" }

func (c *ConfigCommand) Run(args []string) error {
	if err := c.fs.Parse(args); err != nil {
		return err
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--- Interactive Configuration Setup ---")
	
	c.config.AppEnv = c.prompt(reader, "App Environment", c.config.AppEnv)
	c.config.Port = c.prompt(reader, "Port", c.config.Port)
	c.config.APIKey = c.prompt(reader, "API Key (required)", c.config.APIKey)

	fmt.Print("\nSave changes to config.json? (y/n): ")
	confirm, _ := reader.ReadString('\n')
	if strings.ToLower(strings.TrimSpace(confirm)) == "y" {
		if err := c.config.Save(""); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}
		fmt.Println("Configuration saved successfully.")
	} else {
		fmt.Println("Changes discarded.")
	}

	return nil
}

func (c *ConfigCommand) prompt(reader *bufio.Reader, label, current string) string {
	fmt.Printf("%s [%s]: ", label, current)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return current
	}
	return input
}
