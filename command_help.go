package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage: \n\n")

	for name, command := range getCommands() {
		fmt.Printf("%s: %s\n", name, command.description)

	}
	return nil
}
