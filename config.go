package main

import (
	"pokedexcli/internal/pokecache"
)

// config stores application state (pagination API)
type config struct {
	nextURL     *string
	previousURL *string
	currentArea string

	args []string

	Cache   *pokecache.Cache
	Pokedex map[string]Pokemon
}
