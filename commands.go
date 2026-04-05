package main

import (
	"errors"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if _, exists := c.cmds[cmd.Name]; !exists {
		return errors.New("Command doesn't exist")
	}

	if err := c.cmds[cmd.Name](s, cmd); err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}
