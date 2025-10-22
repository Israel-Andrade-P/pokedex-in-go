package pokedex

import (
	"fmt"
	"strconv"
)

type (
	Pokedex struct {
		pokeInfo map[string]Pokemon
	}
	Pokemon struct {
		name string
	}
)

func NewPokedex() Pokedex {
	return Pokedex{pokeInfo: make(map[string]Pokemon, 0)}
}

func (pokedex *Pokedex) RegisterToPokedex(exp int, name string) {
	pokedex.pokeInfo[strconv.Itoa(exp)] = Pokemon{name: name}
}

func (pokedex *Pokedex) PrintPokedex() {
	for key, pokemon := range pokedex.pokeInfo {
		fmt.Printf("%s: %s\n", key, pokemon.name)
	}
}
