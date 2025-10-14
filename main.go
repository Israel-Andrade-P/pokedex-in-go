package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Israel-Andrade-P/pokedex-in-go.git/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

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
				if err := val.callback(); err != nil {
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCommand() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for key, val := range commands {
		fmt.Printf("%s: %s\n", key, val.description)
	}

	return nil
}

func mapCommand() error {
	locations, err := api.GetPokeLocations()
	if err != nil {
		return err
	}
	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}
