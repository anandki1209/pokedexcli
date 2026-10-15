package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {

	resp, err := cfg.pokeapiClient.ListLocations(cfg.Next)
	if err != nil {
		return err
	}

	for _, r := range resp.Results {
		fmt.Println(r.Name)
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocations(cfg.Previous)
	if err != nil {
		return err
	}

	for _, r := range resp.Results {
		fmt.Println(r.Name)
	}
	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	return nil

}
