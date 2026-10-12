package main

import (
	"fmt"

	"github.com/anandki1209/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	if cfg.Next == "" {
		cfg.Next = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	}
	resp, err := pokeapi.GetLocationAreas(cfg.Next)
	if err != nil {
		return err
	}

	for _, r := range resp.Results {
		fmt.Println(r.Name)
	}

	cfg.Next = resp.Next
	if resp.Previous == nil {
		cfg.Previous = ""
	} else {
		cfg.Previous = *resp.Previous
	}

	return nil
}
