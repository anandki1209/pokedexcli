# Pokedex CLI

A small interactive command-line Pokedex written in Go. Browse Pokémon locations, explore the Pokémon found there, try to catch Pokémon, and inspect the ones you have caught. Data is fetched from [PokéAPI](https://pokeapi.co/).

## Requirements

- Go 1.26.3 or later
- Internet access to reach PokéAPI

## Run

From the project directory:

```sh
go run .
```

You can also build a binary with `go build -o pokedex` and run it with `./pokedex`.

## Commands

Enter commands at the `Pokedex >` prompt. Commands are case-insensitive.

| Command | Description |
| --- | --- |
| `help` | Show available commands |
| `map` | Show the next page of location areas |
| `mapb` | Show the previous page of location areas |
| `explore <location>` | List Pokémon found in a location area |
| `catch <pokemon>` | Try to catch a Pokémon |
| `inspect <pokemon>` | Show details for a Pokémon you caught |
| `pokedex` | List your caught Pokémon |
| `exit` | Quit the program |

Example session:

```text
Pokedex > map
canalave-city-area
... 
Pokedex > explore canalave-city-area
tentacool
... 
Pokedex > catch tentacool
Throwing a Pokeball at tentacool...
tentacool was caught!
You may now inspect it with the inspect command.
Pokedex > inspect tentacool
Name: tentacool
...
Pokedex > pokedex
Your Pokedex:
 - tentacool
```

The caught Pokémon list and map page are kept in memory for the current run and reset when you exit. API responses are cached briefly while the program runs.
