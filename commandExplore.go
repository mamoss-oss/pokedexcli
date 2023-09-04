package main

import (
	"encoding/json"
	"fmt"

	"github.com/mamoss-oss/pokedexcli/internal/api"
)

type exploreData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandExplore(c *config, args ...string) error {
	if len(args) != 2 {
		return fmt.Errorf("no region string provided to explore command")
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", args[1])
	body, err := api.CacheOrGet(url, &c.cache)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", args[1])
	data := exploreData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	printPokemons(&data)
	return err
}

func printPokemons(data *exploreData) {
	fmt.Println("Found some Pokemon:")
	for _, pokemon := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
}
