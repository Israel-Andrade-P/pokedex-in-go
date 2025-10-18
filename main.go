package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Israel-Andrade-P/pokedex-in-go.git/api"
	"github.com/Israel-Andrade-P/pokedex-in-go.git/pokecache"
)

type config struct {
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var commands map[string]cliCommand

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
	}

	cfg := &config{
		next:     "https://pokeapi.co/api/v2/location-area?limit=20",
		previous: "",
	}

	cache = pokecache.NewCache(time.Second * 5)

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

		for key, val := range commands {
			if words[0] == key {
				validCommand = true
				if err := val.callback(cfg); err != nil {
					log.Fatalf("Error has occurred: ERR %v", err)
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

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCommand(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for key, val := range commands {
		fmt.Printf("%s: %s\n", key, val.description)
	}

	return nil
}

func mapCommand(cfg *config) error {
	var locationResp api.LocationResponse
	var err error
	cachedData, exists := cache.Get(cfg.next)
	if exists {
		err = json.Unmarshal(cachedData, &locationResp)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		locationResp, err = api.GetPokeLocations(cfg.next)
		if err != nil {
			log.Fatal(err)
		}
		var data []byte
		data, err = json.Marshal(locationResp)
		if err != nil {
			log.Fatal(err)
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

func mapbCommand(cfg *config) error {
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
			log.Fatal(err)
		}
	} else {
		locationResp, err = api.GetPokeLocations(cfg.previous)
		if err != nil {
			log.Fatal(err)
		}
		var data []byte
		data, err = json.Marshal(locationResp)
		if err != nil {
			log.Fatal(err)
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

func getLocationNames(locationResp api.LocationResponse) []string {
	locations := make([]string, 0)

	for _, location := range locationResp.Results {
		locations = append(locations, location.Name)
	}
	return locations
}
