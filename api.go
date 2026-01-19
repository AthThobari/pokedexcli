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
