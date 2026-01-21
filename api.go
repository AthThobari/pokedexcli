package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"pokedexcli/internal/pokecache"
)

// Responses structure from location-area
type locationAreaResponse struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

// Response structure from location-area-detail
type locationAreaDetail struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// fetchLocationAreas fetches location-area data from URL
func fetchLocationAreas(url string, cache *pokecache.Cache) (locationAreaResponse, error) {

	// check cache
	if cached, ok := cache.Get(url); ok {
		fmt.Println("(cache hit)")
		var data locationAreaResponse
		err := json.Unmarshal(cached, &data)
		return data, err
	}

	// fetch from API
	resp, err := http.Get(url)
	if err != nil {
		return locationAreaResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreaResponse{}, err
	}

	// save to cache
	cache.Add(url, body)

	// decode JSON
	var data locationAreaResponse
	err = json.Unmarshal(body, &data)
	return data, nil
}

func fetchLocationAreaDetail(
	name string,
	cache *pokecache.Cache,
) (*locationAreaDetail, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", name)

	if data, ok := cache.Get(url); ok {
		var res locationAreaDetail
		json.Unmarshal(data, &res)
		return &res, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	cache.Add(url, body)

	var result locationAreaDetail
	json.Unmarshal(body, &result)
	return &result, nil
}
