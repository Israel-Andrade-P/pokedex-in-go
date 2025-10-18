package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	PokemonList struct {
		PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
	}
	PokemonEncounter struct {
		Pokemon Pokemon `json:"pokemon"`
	}
	Pokemon struct {
		Name string `json:"name"`
	}
)

func getPokemons(url string) ([]Pokemon, error) {
	var pokemonList PokemonList

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error has occurred: ERR- %v", err)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&pokemonList)
	if err != nil {
		return nil, fmt.Errorf("error has occurred: ERR- %v", err)
	}

	return extractPokemons(pokemonList), nil
}

func extractPokemons(pokeList PokemonList) []Pokemon {
	pokemons := make([]Pokemon, 0)

	for _, encounter := range pokeList.PokemonEncounters {
		pokemons = append(pokemons, encounter.Pokemon)
	}
	return pokemons
}
