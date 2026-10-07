package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {

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
		supportedCommands := getCommands()
		command, ok := supportedCommands[cleanText[0]]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback()
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
	}

	return supportedCommands
}
