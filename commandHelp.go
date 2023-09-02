package main

import "fmt"

func commandHelp(*config) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")

	commands := getCommands()

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
