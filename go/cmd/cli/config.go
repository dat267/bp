package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"bp/config"
	"reflect"
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
	meta := c.config.GetMetadata()
	val := reflect.ValueOf(c.config).Elem()

	for {
		fmt.Println("\n--- rclone-style Configuration Menu ---")
		fmt.Println("v) View current configuration")
		for i, field := range meta {
			fmt.Printf("%d) Edit %s\n", i+1, field.Label)
		}
		fmt.Println("s) Save and Exit")
		fmt.Println("q) Quit without saving")
		fmt.Print("Choose option: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		if choice == "v" {
			c.viewConfig(meta, val)
			continue
		}
		if choice == "s" {
			if err := c.config.Save(""); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
			}
			fmt.Println("Configuration saved.")
			return nil
		}
		if choice == "q" {
			fmt.Println("Exiting without saving.")
			return nil
		}

		idx, err := strconv.Atoi(choice)
		if err == nil && idx > 0 && idx <= len(meta) {
			field := meta[idx-1]
			fVal := val.FieldByName(field.Name)
			
			var validator func(string) error
			if field.Validator != "" {
				validator = config.Validators[field.Validator]
			}

			newValue := c.prompt(reader, field.Label, fVal.String(), validator)
			fVal.SetString(newValue)
		} else {
			fmt.Println("Invalid option.")
		}
	}
}

func (c *ConfigCommand) viewConfig(meta []config.FieldMeta, val reflect.Value) {
	fmt.Printf("\nCurrent Configuration:\n")
	for _, field := range meta {
		fmt.Printf("  %-15s %s\n", field.Label+":", val.FieldByName(field.Name).String())
	}
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
