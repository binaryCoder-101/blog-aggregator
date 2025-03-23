package main

type command struct {
	name string
	args []string
}

type commands struct {
	commandHandler map[string]func(*state, command) error
}
