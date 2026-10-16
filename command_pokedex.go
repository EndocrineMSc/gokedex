package main

import (
	"encoding/json"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	var entries = cfg.cache.GetAll()
	anyCaught := false
	for key := range entries {
		fmt.Println(key)
		var pokemon PokemonDetails
		data, _ := cfg.cache.Get(key)
		err := json.Unmarshal(data, &pokemon)
		if err == nil && pokemon.IsCaught {
			anyCaught = true
			fmt.Printf("- %s\n", pokemon.Name)
		}
	}
	if !anyCaught {
		fmt.Println("No pokemon caught yet!")
	}

	return nil
}
