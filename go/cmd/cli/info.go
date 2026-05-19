package main

import (
	"flag"
	"fmt"

	"bp/config"
)

// InfoCommand displays environment information
type InfoCommand struct {
	fs     *flag.FlagSet
	config *config.Config
}

func NewInfoCommand(cfg *config.Config) *InfoCommand {
	return &InfoCommand{
		fs:     flag.NewFlagSet("info", flag.ContinueOnError),
		config: cfg,
	}
}

func (c *InfoCommand) Name() string        { return c.fs.Name() }
func (c *InfoCommand) Description() string { return "Displays environment information" }

func (c *InfoCommand) Run(args []string) error {
	if err := c.fs.Parse(args); err != nil {
		return err
	}
	fmt.Printf("Environment: %s\n", c.config.AppEnv)
	fmt.Printf("Port:        %s\n", c.config.Port)
	fmt.Printf("Verbose:     %v\n", c.config.Verbose)
	return nil
}
