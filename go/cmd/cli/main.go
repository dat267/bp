package main

import (
	"flag"
	"fmt"
	"os"

	"bp/config"
)

// Command defines the interface for CLI commands
type Command interface {
	Run(args []string) error
	Name() string
	Description() string
}

func main() {
	cfg := config.LoadConfig()

	commands := []Command{
		NewHelloCommand(cfg),
		NewInfoCommand(cfg),
	}

	usage := func() {
		fmt.Printf("Usage: %s <command> [arguments]\n", os.Args[0])
		fmt.Println("\nAvailable commands:")
		for _, cmd := range commands {
			fmt.Printf("  %-10s %s\n", cmd.Name(), cmd.Description())
		}
		fmt.Println("\nUse \"<command> --help\" for more information about a command.")
	}

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	subcommand := os.Args[1]
	if subcommand == "help" || subcommand == "-h" || subcommand == "--help" {
		usage()
		return
	}

	for _, cmd := range commands {
		if cmd.Name() == subcommand {
			if err := cmd.Run(os.Args[2:]); err != nil {
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
