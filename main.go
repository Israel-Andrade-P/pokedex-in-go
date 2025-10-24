package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Israel-Andrade-P/pokedex-in-go.git/cmds"
)

func main() {
	commands := cmds.GetCmds()
	cfg := cmds.GetConfig()

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
				if err := val.Callback(cfg, parameter); err != nil {
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
