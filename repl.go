package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	for {
		// Create a new scanner for reading from standard input
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")

		// Scan the user input and store it in a variable
		scanner.Scan()
		userCommand := scanner.Text()
		userCommand = cleanText(userCommand)
		commands := getCommands()
		c, ok := commands[userCommand]
		if !ok {
			fmt.Println("Sorry, command not found. Try 'help' for usage guidelines.")
			continue
		}
		_ = c.callback()
	}
}

func cleanText(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	return s
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
