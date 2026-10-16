package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/anandki1209/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient     pokeapi.Client
	supportedCommands map[string]cliCommand
	UserPokedex       map[string]pokeapi.Pokemon
	Next              *string
	Previous          *string
}

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		cleanText := cleanInput(input)

		if len(cleanText) == 0 { // guarding against index out or error
			continue
		}
		supportedCommands := cfg.supportedCommands
		command, ok := supportedCommands[cleanText[0]]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg, cleanText[1:]...)
		if err != nil {
			fmt.Print(err)
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error trying to read input", err)
	}
}

func cleanInput(text string) []string {
	// lowercase string and trim any leading or trailing whitespace
	// split string on space
	lowercaseText := strings.ToLower(text)

	result := strings.Fields(lowercaseText)
	return result

}

func getCommands() map[string]cliCommand {
	supportedCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "explore the pokemon for a given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "tries to catch the pokemon",
			callback:    commandCatch,
		},
	}
	return supportedCommands
}
