package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []PokeLocation `json:"results"`
}
type PokeLocation struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func GetPokeLocations(url string) (LocationResponse, error) {
	var locations LocationResponse

	res, err := http.Get(url)
	if err != nil {
		return locations, fmt.Errorf("error has occurred: ERR- %v", err)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&locations)
	if err != nil {
		return locations, fmt.Errorf("error has occurred: ERR- %v", err)
	}

	return locations, nil
}
