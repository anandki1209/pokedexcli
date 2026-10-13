package main

import (
	"time"

	"github.com/anandki1209/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := config{
		pokeapiClient:     pokeapi.NewClient(5 * time.Second),
		supportedCommands: getCommands(),
	}
	startRepl(&cfg)

}
