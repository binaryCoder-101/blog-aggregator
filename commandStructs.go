package main

import "fmt"

type command struct {
	commandName string
	commandArgs []string
}

type commands struct {
	commandHandler map[string]func(*state, command) error
}

func (c *commands) register(commandName string, handler func(*state, command) error) {
	c.commandHandler[commandName] = handler
}

func (c *commands) run(s *state, cmd command) error {
	if s == nil {
		return fmt.Errorf("error: state is uninitialized")
	}
	c.commandHandler[cmd.commandName](s, cmd)
	return nil
}
