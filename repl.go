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
	callback    func(*config) error
}

type config struct {
	next     string
	previous string
}

func startRepl() {

	userConfig := config{
		next: "https://pokeapi.co/api/v2/location-area/",
	}

	for {
		// Create a new scanner for reading from standard input
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")

		// Scan the user input and store it in a variable
		scanner.Scan()
		userCommand := scanner.Text()
		userCommand = CleanText(userCommand)
		commands := getCommands()
		c, ok := commands[userCommand]
		if !ok {
			fmt.Println("Sorry, command not found. Try 'help' for usage guidelines.")
			continue
		}
		err := c.callback(&userConfig)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func CleanText(s string) string {
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
		"map": {
			name:        "map",
			description: "Show the next location area",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous location area",
			callback:    commandMapb,
		},
	}
}
