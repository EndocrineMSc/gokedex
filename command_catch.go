package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("name a pokemon")
	}

	pokemonName := args[0]
	data, exists := cfg.cache.Get(pokemonName)
	if exists {
		var details PokemonDetails
		err := json.Unmarshal(data, &details)
		if err != nil {
			return fmt.Errorf("could not unmarshal pokemon: %w", err)
		}
		if details.IsCaught {
			fmt.Println(details.Name + " was already caught!")
			return nil
		}
		if tryCatch(details) {
			details.IsCaught = true
			deets, err := json.Marshal(details)
			if err != nil {
				return fmt.Errorf("could not marshal pokemon: %w", err)
			}
			cfg.cache.Add(pokemonName, deets)
		}
		return nil
	}

	res, err := http.Get(cfg.apiBaseUrl + cfg.pokeUrl + pokemonName)
	if err != nil {
		return fmt.Errorf("could not fetch pokemon: %w", err)
	}
	defer res.Body.Close()
	var details PokemonDetails
	err = json.NewDecoder(res.Body).Decode(&details)
	if err != nil {
		return fmt.Errorf("could not decode pokemon: %w", err)
	}
	if details.Name == "" {
		return errors.New("pokemon not found")
	}
	isCaught := tryCatch(details)
	if isCaught {
		details.IsCaught = true
	}
	deets, err := json.Marshal(details)
	if err != nil {
		return fmt.Errorf("could not marshal pokemon: %w", err)
	}
	cfg.cache.Add(pokemonName, deets)
	return nil
}

func tryCatch(pd PokemonDetails) bool {
	name := pd.Name
	fmt.Println("Throwing a Pokeball at " + name + "...")
	chance := 100.0 / (float64(pd.Exp) / 100.0)
	if chance <= 5 {
		chance = 5
	} else if chance > 95 {
		chance = 95
	}
	wiggle := rand.Float64() * 100.0
	success := wiggle <= chance
	if success {
		fmt.Println(name + " was caught!")
	} else {
		fmt.Println(name + " escaped!")
	}
	return success
}
