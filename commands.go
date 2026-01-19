package main

// cliCommand represents a single command in the REPL
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}
