package main

import "fmt"

func commandPokedex(c *config, args ...string) error {
	if len(c.pokedex) == 0 {
		return fmt.Errorf("cannot display empty pokedex")
	}
	fmt.Println("Your Pokedex:")
	for key := range c.pokedex {
		fmt.Printf(" - %s\n", key)
	}
	return nil
}
