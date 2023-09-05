package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mamoss-oss/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	cache    pokecache.Cache
	next     string
	previous string
	pokedex  map[string]PokemonData
}

func startRepl() {

	userConfig := config{
		cache:   pokecache.NewCache(time.Second * 300),
		next:    "https://pokeapi.co/api/v2/location-area/",
		pokedex: make(map[string]PokemonData),
	}

	for {
		// Create a new scanner for reading from standard input
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")

		// Scan the user input and store it in a variable
		scanner.Scan()
		userCommand := scanner.Text()
		userCommand = CleanText(userCommand)
		split := strings.Split(userCommand, " ")
		commands := getCommands()
		c, ok := commands[split[0]]
		if !ok {
			fmt.Println("Sorry, command not found. Try 'help' for usage guidelines.")
			continue
		}
		err := c.callback(&userConfig, split...)
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
		"explore": {
			name:        "explore",
			description: "Explore the Pokemons in a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Check out a known pokemon in your pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all Pokemons in your current pokedex",
			callback:    commandPokedex,
		},
	}
}
