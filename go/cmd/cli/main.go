package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"bp/config"
)

// Command defines the interface for CLI commands
type Command interface {
	Run(args []string) error
	Name() string
	Description() string
}

func main() {
	// 1. Calculate default config path (portable: same folder as executable)
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exDir := filepath.Dir(ex)
	defaultConfigPath := filepath.Join(exDir, "config.json")

	// 2. Preliminary parse for global flags (like --config, --verbose)
	var configPath string
	var verbose bool
	var args []string
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if (arg == "--config" || arg == "-c") && i+1 < len(os.Args) {
			configPath = os.Args[i+1]
			i++ // Skip the value
		} else if arg == "--verbose" || arg == "-v" {
			verbose = true
		} else {
			args = append(args, arg)
		}
	}

	if configPath == "" {
		configPath = defaultConfigPath
	}

	cfg := config.LoadConfig(configPath)
	if verbose {
		cfg.Verbose = true
	}

	commands := []Command{
		NewHelloCommand(cfg),
		NewInfoCommand(cfg),
		NewConfigCommand(cfg),
	}

	usage := func() {
		fmt.Printf("Usage: %s [global flags] <command> [arguments]\n", os.Args[0])
		fmt.Println("\nGlobal flags:")
		fmt.Println("  --config, -c <path>  Path to config file")
		fmt.Println("\nAvailable commands:")
		for _, cmd := range commands {
			fmt.Printf("  %-10s %s\n", cmd.Name(), cmd.Description())
		}
		fmt.Println("\nUse \"<command> --help\" for more information about a command.")
	}

	if len(args) < 1 {
		usage()
		os.Exit(1)
	}

	subcommand := args[0]
	if subcommand == "help" || subcommand == "-h" || subcommand == "--help" {
		usage()
		return
	}

	for _, cmd := range commands {
		if cmd.Name() == subcommand {
			if err := cmd.Run(args[1:]); err != nil {
				if err != flag.ErrHelp {
					fmt.Fprintf(os.Stderr, "Error: %v\n", err)
					os.Exit(1)
				}
				return
			}
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown command: %s\n", subcommand)
	usage()
	os.Exit(1)
}
