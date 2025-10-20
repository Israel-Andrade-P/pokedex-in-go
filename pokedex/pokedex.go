package pokedex

import "fmt"

type (
	Pokedex struct {
		pokeInfo map[string]Pokemon
	}
	Pokemon struct {
		name string
	}
)

func newPokedex() Pokedex {
	return Pokedex{pokeInfo: make(map[string]Pokemon, 0)}
}

func printPokedex() {
	myPokedex := newPokedex()
	for key, pokemon := range myPokedex.pokeInfo {
		fmt.Printf("%s: %s", key, pokemon.name)
	}
}
