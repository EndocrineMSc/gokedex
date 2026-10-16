package main

type PokemonDetails struct {
	Id       int         `json:"id"`
	Name     string      `json:"name"`
	Exp      int         `json:"base_experience"`
	Height   int         `json:"height"`
	Weight   int         `json:"weight"`
	Stats    []Stat      `json:"stats"`
	Types    []PokeTypes `json:"types"`
	IsCaught bool
}

type Location struct {
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
}

type LocationAreaResponse struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}

type Stat struct {
	Val  int `json:"base_stat"`
	Name struct {
		Val string `json:"name"`
	} `json:"stat"`
}

type PokeTypes struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}
