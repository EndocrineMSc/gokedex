package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("location required")
	}

	locName := args[0]
	data, exists := cfg.cache.Get(locName)

	if exists {
		var pokemon []string
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return fmt.Errorf("could not unmarshal pokemon: %w", err)
		}
		for _, name := range pokemon {
			fmt.Println(name)
		}
		return nil
	}

	fullUrl := cfg.apiBaseUrl + cfg.locUrl + locName
	res, err := http.Get(fullUrl)
	if err != nil {
		return fmt.Errorf("could not fetch location-area: %w", err)
	}
	defer res.Body.Close()
	var location Location
	err = json.NewDecoder(res.Body).Decode(&location)
	if err != nil {
		return fmt.Errorf("could not parse location-area: %w", err)
	}

	encounters := location.PokemonEncounters

	var names []string
	for _, encounter := range encounters {
		fmt.Println(encounter.Pokemon.Name)
		names = append(names, encounter.Pokemon.Name)
	}

	cachePokemons, err := json.Marshal(names)
	if err != nil {
		return fmt.Errorf("could not marshal pokemon: %w", err)
	}
	cfg.cache.Add(locName, cachePokemons)

	return nil
}
