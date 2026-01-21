package main

import (
	"fmt"
	"os"
)

// exit command
func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// help command (using closure)

func commandHelp(commands map[string]cliCommand) func(*config) error {
	return func(cfg *config) error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		fmt.Println()

		for _, cmd := range commands {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		}
		return nil
	}
}

// map command (next page)
func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.nextURL != nil {
		url = *cfg.nextURL
	}

	data, err := fetchLocationAreas(url, cfg.Cache)
	if err != nil {
		return err
	}

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	cfg.nextURL = data.Next
	cfg.previousURL = data.Previous
	return nil
}

// mapb command (previous page)
func commandMapBack(cfg *config) error {
	if cfg.previousURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	data, err := fetchLocationAreas(*cfg.previousURL, cfg.Cache)
	if err != nil {
		return err
	}

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	cfg.nextURL = data.Next
	cfg.previousURL = data.Previous
	return nil
}

func commandExplore(cfg *config) error {
if cfg.currentArea == "" {
return fmt.Errorf("you must provide a location area")
}

fmt.Printf("Exploring %s...\n", cfg.currentArea)
fmt.Println("Found Pokemon:")

res, err := fetchLocationAreaDetail(cfg.currentArea, cfg.Cache)
if err != nil {
return err
}

for _, p := range res.PokemonEncounters {
fmt.Printf("- %s\n", p.Pokemon.Name)
}

return nil
}
