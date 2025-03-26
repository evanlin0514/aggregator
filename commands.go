package main

import "fmt"

type command struct {
    name string
    args []string
}

type commands struct {
    Handlers map[string]func(*state, command)error 
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.Handlers[name] = f
}

func (c *commands) run (s *state, cmd command) error{
	f, ok := c.Handlers[cmd.name]
	if !ok {
		return fmt.Errorf("handler not exist")
	}

	return f(s, cmd)
}