package main

import (
	"flag"
	"fmt"
	"os"
)

// Command defines the interface for CLI commands
type Command interface {
	Run(args []string) error
	Name() string
	Description() string
}

// HelloCommand is an example implementation of a command
type HelloCommand struct {
	fs   *flag.FlagSet
	name string
}

func NewHelloCommand() *HelloCommand {
	hc := &HelloCommand{
		fs: flag.NewFlagSet("hello", flag.ContinueOnError),
	}
	hc.fs.StringVar(&hc.name, "name", "World", "name to greet")
	return hc
}

func (c *HelloCommand) Name() string        { return c.fs.Name() }
func (c *HelloCommand) Description() string { return "Greets a person" }

func (c *HelloCommand) Run(args []string) error {
	if err := c.fs.Parse(args); err != nil {
		return err
	}
	fmt.Printf("Hello, %s!\n", c.name)
	return nil
}

func main() {
	commands := []Command{
		NewHelloCommand(),
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [arguments]")
		fmt.Println("\nAvailable commands:")
		for _, cmd := range commands {
			fmt.Printf("  %-10s %s\n", cmd.Name(), cmd.Description())
		}
		os.Exit(1)
	}

	subcommand := os.Args[1]
	for _, cmd := range commands {
		if cmd.Name() == subcommand {
			if err := cmd.Run(os.Args[2:]); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown command: %s\n", subcommand)
	os.Exit(1)
}
