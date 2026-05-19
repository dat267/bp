package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
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

	for {
		fmt.Println("\n--- rclone-style Configuration Menu ---")
		fmt.Println("1) View current configuration")
		fmt.Println("2) Edit App Environment")
		fmt.Println("3) Edit Port")
		fmt.Println("4) Edit API Key")
		fmt.Println("s) Save and Exit")
		fmt.Println("q) Quit without saving")
		fmt.Print("Choose option: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			c.viewConfig()
		case "2":
			c.config.AppEnv = c.prompt(reader, "App Environment", c.config.AppEnv, nil)
		case "3":
			c.config.Port = c.prompt(reader, "Port", c.config.Port, validatePort)
		case "4":
			c.config.APIKey = c.prompt(reader, "API Key", c.config.APIKey, validateNotEmpty)
		case "s":
			if err := c.config.Save(""); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
			}
			fmt.Println("Configuration saved.")
			return nil
		case "q":
			fmt.Println("Exiting without saving.")
			return nil
		default:
			fmt.Println("Invalid option.")
		}
	}
}

func (c *ConfigCommand) viewConfig() {
	fmt.Printf("\nCurrent Configuration:\n")
	fmt.Printf("  AppEnv:  %s\n", c.config.AppEnv)
	fmt.Printf("  Port:    %s\n", c.config.Port)
	fmt.Printf("  APIKey:  %s\n", c.config.APIKey)
}

func (c *ConfigCommand) prompt(reader *bufio.Reader, label, current string, validator func(string) error) string {
	for {
		fmt.Printf("%s [%s]: ", label, current)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			return current
		}

		if validator != nil {
			if err := validator(input); err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
		}
		return input
	}
}

func validatePort(val string) error {
	port, err := strconv.Atoi(val)
	if err != nil {
		return fmt.Errorf("port must be a number")
	}
	if port < 1 || port > 65535 {
		return fmt.Errorf("port must be between 1 and 65535")
	}
	return nil
}

func validateNotEmpty(val string) error {
	if strings.TrimSpace(val) == "" {
		return fmt.Errorf("value cannot be empty")
	}
	return nil
}
