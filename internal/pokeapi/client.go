package pokeapi

import (
	"encoding/json"
	"net/http"
)

type LocationAreaResponse struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func FetchLocationAreas(url string) (LocationAreaResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.close()

	var data LocationAreaResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	return data, err
}
