package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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
	if len(cfg.args) == 0 {
		return fmt.Errorf("you must provide a location area")
	}

	area := cfg.args[0]
	cfg.currentArea = area

	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")

	res, err := fetchLocationAreaDetail(area, cfg.Cache)
	if err != nil {
		return err
	}

	for _, p := range res.PokemonEncounters {
		fmt.Printf("- %s\n", p.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config) error {
	// argument validation
	if len(cfg.args) == 0 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	name := strings.ToLower(cfg.args[0])

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	// fetch pokemon
	pokemon, err := fetchPokemon(name, cfg.Cache)
	if err != nil {
		return err
	}

	// seed random (necessary)
	rand.Seed(time.Now().UnixNano())

	// logic chance
	chance := rand.Intn(pokemon.BaseExperience + 1)

	if chance > pokemon.BaseExperience/2 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	// successfully captured
	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.Pokedex[pokemon.Name] = *pokemon

	return nil
}
