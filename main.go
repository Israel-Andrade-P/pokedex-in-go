package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Israel-Andrade-P/pokedex-in-go.git/api"
	"github.com/Israel-Andrade-P/pokedex-in-go.git/pokecache"
	"github.com/Israel-Andrade-P/pokedex-in-go.git/pokedex"
)

type config struct {
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

var commands map[string]cliCommand

var myPokedex pokedex.Pokedex

var cache *pokecache.Cache

func main() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCommand,
		},
		"map": {
			name:        "map",
			description: "Displays available locations",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous locations page",
			callback:    mapbCommand,
		},
		"explore": {
			name:        "explore",
			description: "Lists pokemon names of a certain location. Accepts a parameter location name.\nEx: explore <location>",
			callback:    exploreCommand,
		},
		"catch": {
			name:        "catch",
			description: "Attempts catching a specific Pokemon.\nEx: catch <pokemon>",
			callback:    catchCommand,
		},
		"inspect": {
			name:        "inspect",
			description: "It will show information about a pokemon if caught it before.\nEx: inspect <pokemon>",
			callback:    inspectCommand,
		},
	}

	cfg := &config{
		next:     "https://pokeapi.co/api/v2/location-area?limit=20",
		previous: "",
	}

	myPokedex = pokedex.NewPokedex()

	cache = pokecache.NewCache(time.Second * 10)

	sc := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Pokedex > ")
		scanned := sc.Scan()
		if !scanned {
			break
		}

		input := sc.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}

		validCommand := false
		var parameter string

		for key, val := range commands {
			if words[0] == key {
				validCommand = true
				if len(words) > 1 {
					parameter = cleanParameter(words[1:])
				}
				if err := val.callback(cfg, parameter); err != nil {
					log.Fatal(err)
				}
			}
		}
		if !validCommand {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func cleanParameter(words []string) string {
	if len(words) == 1 {
		return strings.ToLower(words[0])
	}
	return strings.ToLower(strings.Join(words, "-"))
}

func commandExit(cfg *config, parameter string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCommand(cfg *config, parameter string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for key, val := range commands {
		fmt.Printf("%s: %s\n", key, val.description)
	}

	return nil
}

func mapCommand(cfg *config, parameter string) error {
	var locationResp api.LocationResponse
	var err error
	cachedData, exists := cache.Get(cfg.next)
	if exists {
		err = json.Unmarshal(cachedData, &locationResp)
		if err != nil {
			return err
		}
	} else {
		locationResp, err = api.GetPokeLocations(cfg.next)
		if err != nil {
			return err
		}
		var data []byte
		data, err = json.Marshal(locationResp)
		if err != nil {
			return err
		}
		cache.Add(cfg.next, data)
	}

	for _, location := range getLocationNames(locationResp) {
		fmt.Println(location)
	}

	cfg.next = locationResp.Next
	cfg.previous = locationResp.Previous

	return nil
}

func mapbCommand(cfg *config, parameter string) error {
	if cfg.previous == "" {
		fmt.Println("you're on the first page.")
		return nil
	}
	var locationResp api.LocationResponse
	var err error
	cachedData, exists := cache.Get(cfg.previous)
	if exists {
		err = json.Unmarshal(cachedData, &locationResp)
		if err != nil {
			return err
		}
	} else {
		locationResp, err = api.GetPokeLocations(cfg.previous)
		if err != nil {
			return err
		}
		var data []byte
		data, err = json.Marshal(locationResp)
		if err != nil {
			return err
		}
		cache.Add(cfg.previous, data)
	}

	for _, location := range getLocationNames(locationResp) {
		fmt.Println(location)
	}

	cfg.next = locationResp.Next
	cfg.previous = locationResp.Previous

	return nil
}

func exploreCommand(cfg *config, parameter string) error {
	var pokeNames []string
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", parameter)
	cachedData, exists := cache.Get(url)
	if exists {
		err := json.Unmarshal(cachedData, &pokeNames)
		if err != nil {
			return err
		}
	} else {
		var err error
		pokeNames, err = api.GetPokemons(url)
		if err != nil {
			return err
		}
		data, err := json.Marshal(pokeNames)
		if err != nil {
			return err
		}
		cache.Add(url, data)
	}

	for _, name := range pokeNames {
		fmt.Println(name)
	}
	return nil
}

func catchCommand(cfg *config, parameter string) error {
	fmt.Printf("Throwing a pokeball at %s\n", parameter)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", parameter)
	cachedData, exists := cache.Get(url)
	var pokeInfo api.PokeInfo
	if exists {
		if err := json.Unmarshal(cachedData, &pokeInfo); err != nil {
			return err
		}
	} else {
		var err error
		pokeInfo, err = api.GetPokeInfo(url)
		if err != nil {
			return err
		}
		data, err := json.Marshal(pokeInfo)
		if err != nil {
			return err
		}
		cache.Add(url, data)
	}
	prob := catchOrFail(pokeInfo.BaseExp)
	if prob > 40 {
		fmt.Printf("%s escaped!\n", parameter)
	} else {
		fmt.Printf("%s was caught!\n", parameter)
		stats := make(map[string]int)
		types := make([]string, 0)
		for _, statInfo := range pokeInfo.Stats {
			stats[statInfo.Stat.Name] = statInfo.BaseStat
		}
		for _, typeInfo := range pokeInfo.Types {
			types = append(types, typeInfo.Type.Name)
		}
		pokemon := pokedex.CreatePokemon(parameter, pokeInfo.Height, pokeInfo.Weight, stats, types)
		myPokedex.RegisterToPokedex(parameter, pokemon)
	}
	return nil
}

func inspectCommand(cfg *config, parameter string) error {
	myPokedex.InspectPokemon(parameter)
	return nil
}

func catchOrFail(baseExp int) int {
	return (rand.Intn(baseExp) + 1)
}

func getLocationNames(locationResp api.LocationResponse) []string {
	locations := make([]string, 0)

	for _, location := range locationResp.Results {
		locations = append(locations, location.Name)
	}
	return locations
}
