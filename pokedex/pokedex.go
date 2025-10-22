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
	if pokemon, exists := pokedex.pokeInfo[name]; exists {
		pokePrint(pokemon)
		return
	}
	fmt.Println("you have not caught that pokemon")
}

func pokePrint(pokemon Pokemon) {
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	fmt.Printf("Stats: ")
	for key, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d", key, stat)
	}

	for _, pokeType := range pokemon.Types {
		fmt.Printf("- %s", pokeType)
	}
}
