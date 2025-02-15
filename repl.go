package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/bootdotdev/pokedexcli/internal/pokeapi"
)
type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	areaName *string
	pokedex map[string]pokeapi.Pokemon
	
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		if len(words) > 1 {
			cfg.areaName=&words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config,...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
            description: "Explore a specific location",
            callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
            description: "Catch a specific Pokémon",
            callback:    commandCatch,
		},
		"inspect":{
			name:        "inspect",
            description: "Inspect a specific Pokémon",
            callback:    commandInspect,
		},
		"pokedex":{
			name:        "pokedex",
            description: "View your Pokédex",
            callback:    commandPokedex,
		},
	}
}
