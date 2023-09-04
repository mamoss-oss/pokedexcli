package main

import (
	"encoding/json"
	"fmt"

	"github.com/mamoss-oss/pokedexcli/internal/api"
)

type AreaData struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(c *config, args ...string) error {
	body, err := api.CacheOrGet(c.next, &c.cache)
	if err != nil {
		return err
	}
	areas := AreaData{}
	err = json.Unmarshal(body, &areas)
	if err != nil {
		return err
	}
	updateConf(c, &areas)
	printAreas(&areas)
	return err
}

func commandMapb(c *config, args ...string) error {
	if c.previous == "" {
		fmt.Println("Previous is empty")
		return nil
	}
	body, err := api.CacheOrGet(c.previous, &c.cache)
	if err != nil {
		return err
	}
	areas := AreaData{}
	err = json.Unmarshal(body, &areas)
	if err != nil {
		return err
	}
	updateConf(c, &areas)
	printAreas(&areas)
	return err
}

// updateConf accepts references to config and AreaData structs.
// It validates if the received data for previous or next is nil,
// before updating the config struct.
func updateConf(c *config, areas *AreaData) {
	if areas.Previous != nil {
		c.previous = *areas.Previous
	}
	if areas.Next != nil {
		c.next = *areas.Next
	}
}

// printAreas accepts a reference to an AreaData struct.
// It print ranges over the result names.
func printAreas(areas *AreaData) {
	for _, area := range areas.Results {
		println(area.Name)
	}
}
