package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(cfg *config, args ...string) error {
	return printLocations(false, cfg)
}

func commandMapb(cfg *config, args ...string) error {
	return printLocations(true, cfg)
}

func printLocations(isBack bool, cfg *config) error {
	limit := 20
	cache := cfg.cache

	if cfg.currentOffset == 0 && isBack {
		fmt.Println("you're on the first page")
		return nil
	} else if isBack {
		cfg.currentOffset -= limit
	} else if cfg.hasLeftZero {
		cfg.currentOffset += limit
	}

	query := fmt.Sprintf("?offset=%d&limit=%d", cfg.currentOffset, limit)
	fullUrl := cfg.apiBaseUrl + cfg.locUrl + query

	if !cfg.hasLeftZero {
		cfg.currentOffset += limit
		cfg.hasLeftZero = true
	}

	hit, exists := cache.Get(fullUrl)
	if exists {
		var hits []string
		err := json.Unmarshal(hit, &hits)
		if err != nil {
			return fmt.Errorf("could not parse hit from cache: %w", err)
		}
		for _, v := range hits {
			fmt.Printf("%v\n", v)
		}
		return nil
	}

	res, err := http.Get(fullUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var body LocationAreaResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&body); err != nil {
		return err
	}
	var names []string
	for _, v := range body.Results {
		fmt.Println(v.Name)
		names = append(names, v.Name)
	}
	data, err := json.Marshal(names)
	if err != nil {
		return fmt.Errorf("could not marshal names: %w", err)
	}
	cache.Add(fullUrl, data)
	return nil
}
