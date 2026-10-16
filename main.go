package main

import (
	"time"

	"github.com/anandki1209/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := config{
		pokeapiClient: pokeClient,
		UserPokedex:   make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)

}
