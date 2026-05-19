package main

import (
	"flag"
	"fmt"

	"bp/config"
)

// HelloCommand is an example implementation of a command
type HelloCommand struct {
	fs     *flag.FlagSet
	name   string
	config *config.Config
}

func NewHelloCommand(cfg *config.Config) *HelloCommand {
	hc := &HelloCommand{
		fs:     flag.NewFlagSet("hello", flag.ContinueOnError),
		config: cfg,
	}
	// Use config (which includes file/env fallbacks) as the default value
	defaultName := "World"
	if cfg.AppEnv == "production" {
		defaultName = "Production User"
	}
	hc.fs.StringVar(&hc.name, "name", defaultName, "name to greet")
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
