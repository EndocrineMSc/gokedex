package main

import (
	"time"

	"github.com/EndocrineMSc/pokedex/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *config, args ...string) error
}

func main() {
	cfg := &config{
		commands:      getCommands(),
		apiBaseUrl:    "https://pokeapi.co/api/v2/",
		locUrl:        "location-area/",
		pokeUrl:       "pokemon/",
		currentOffset: 0,
		cache:         internal.NewCache(5 * time.Second),
	}
	startRepl(cfg)
}
