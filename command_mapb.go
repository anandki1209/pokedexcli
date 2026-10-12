package main

import (
	"fmt"

	"github.com/anandki1209/pokedexcli/internal/pokeapi"
)

func commandMapb(cfg *config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	resp, err := pokeapi.GetLocationAreas(cfg.Previous)
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
