package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("type a pokemon to inspect after the command")
	}

	data, exists := cfg.cache.Get(args[0])
	if !exists {
		return errors.New("pokemon does not seen yet")
	}
	var pokemon PokemonDetails
	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return fmt.Errorf("could not parse pokemon data: %w", err)
	}
	if !pokemon.IsCaught {
		return errors.New("pokemon not caught")
	}

	fmt.Println("Name: " + pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		name := stat.Name.Val
		val := stat.Val
		fmt.Printf("    -%s: %d\n", name, val)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("    -%s\n", typ.Type.Name)
	}
	fmt.Println("IsCaught:", pokemon.IsCaught)

	return nil
}
