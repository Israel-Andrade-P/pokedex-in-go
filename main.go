package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Israel-Andrade-P/pokedex-in-go.git/api"
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

var locationResp api.LocationResponse

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

	url := "https://pokeapi.co/api/v2/location/"
	var err error

	locationResp, err = api.GetPokeLocations(url)
	if err != nil {
		log.Fatal(err)
	}

	cfg := &config{
		next:     locationResp.Next,
		previous: locationResp.Previous,
	}

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
	for _, location := range getLocationNames(locationResp) {
		fmt.Println(location)
	}
	var err error
	locationResp, err = api.GetPokeLocations(cfg.next)
	if err != nil {
		return err
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
	var err error
	locationResp, err = api.GetPokeLocations(cfg.previous)
	if err != nil {
		return err
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
