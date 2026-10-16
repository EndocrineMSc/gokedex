package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/EndocrineMSc/pokedex/internal"
)

type config struct {
	commands      map[string]cliCommand
	apiBaseUrl    string
	locUrl        string
	pokeUrl       string
	currentOffset int
	hasLeftZero   bool
	cache         *internal.Cache
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanInput := cleanInput(input)
		if len(cleanInput) == 0 {
			continue
		}

		commandName := cleanInput[0]
		command, exists := cfg.commands[commandName]
		if exists {
			err := command.callback(cfg, cleanInput[1:]...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command.")
			continue
		}
	}
}

// cleanInput will return a slice of whitespace-trimmed strings
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get 20 locations from the pokedex - incremental",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get last 20 locations from the pokedex",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location - example: explore vertania-city",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon - example: catch pikachu",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught pokemon - example: inspect pikachu",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show all caught pokemon",
			callback:    commandPokedex,
		},
	}
}
