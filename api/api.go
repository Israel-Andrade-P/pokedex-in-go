package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []pokeLocation `json:"results"`
}
type pokeLocation struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func GetPokeLocations() ([]string, error) {
	url := "https://pokeapi.co/api/v2/location/"
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error has occurred: ERR- %v", err)
	}

	defer res.Body.Close()

	var locations locationResponse

	err = json.NewDecoder(res.Body).Decode(&locations)
	if err != nil {
		return nil, fmt.Errorf("Error has occurred: ERR- %v", err)
	}

	return getLocationNames(locations), nil
}

func getLocationNames(locationResp locationResponse) []string {
	locations := make([]string, 0)

	for _, location := range locationResp.Results {
		locations = append(locations, location.Name)
	}
	return locations
}
