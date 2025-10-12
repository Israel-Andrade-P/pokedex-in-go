package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(cleanInput("    hello     WORLd  "))
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
