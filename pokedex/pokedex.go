package pokedex

import (
	"fmt"
)

type (
	Pokedex struct {
		pokeInfo map[string]Pokemon
	}
	Pokemon struct {
		Name   string
		Height int
		Weight int
		Stats  map[string]int
		Types  []string
	}
)

func NewPokedex() Pokedex {
	return Pokedex{pokeInfo: make(map[string]Pokemon, 0)}
}

func (pokedex *Pokedex) RegisterToPokedex(name string, pokemon Pokemon) {
	pokedex.pokeInfo[name] = pokemon
}

func (pokedex *Pokedex) InspectPokemon(name string) {
	if _, exists := pokedex.pokeInfo[name]; exists {
		pokedex.pokePrint(name)
		return
	}
	fmt.Println("you have not caught that pokemon")
}

func CreatePokemon(name string, height, weight int, stats map[string]int, types []string) Pokemon {
	return Pokemon{
		Name:   name,
		Height: height,
		Weight: weight,
		Stats:  stats,
		Types:  types,
	}
}

func (pokedex *Pokedex) pokePrint(pokeName string) {
	pokemon := pokedex.pokeInfo[pokeName]
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	fmt.Println("Stats:")
	for key, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d\n", key, stat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("- %s\n", pokeType)
	}
}
