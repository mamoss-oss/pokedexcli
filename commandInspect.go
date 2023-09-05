package main

import "fmt"

func commandInspect(c *config, args ...string) error {
	if len(args) != 2 {
		return fmt.Errorf("no pokemon string provided to inspect command")
	}
	pokemon, ok := c.pokedex[args[1]]
	if !ok {
		return fmt.Errorf("%s not found in pokedex", args[1])
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Println("Stats:")
	for _, stats := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokemonType.Type.Name)
	}
	return nil
}
