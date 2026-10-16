package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("you must provide a  pokemon name")
	}

	pokemonName := args[0]

	pokemon, err := c.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	c.UserPokedex[pokemon.Name] = pokemon

	return nil

}
