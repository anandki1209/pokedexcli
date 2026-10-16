package main

import "fmt"

func commandPokedex(c *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for name := range c.UserPokedex {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
