package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage: \n\n")

	supportedCommands := getCommands()
	for name, command := range supportedCommands {
		fmt.Printf("%s: %s\n", name, command.description)

	}
	return nil
}
