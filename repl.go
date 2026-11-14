package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/vitasukka/pokedex/pokecache"
)

type config struct {
	supportedCommands map[string]cliCommand
	Previous          string
	Next              string
	PokeCache         pokecache.Cache
	PokeDex           map[string]Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, arg string) error
}

var cliCommands = map[string]cliCommand{
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
		description: "shows you the next 20 areas on the map",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "shows you the last 20 areas on the map you have visited",
		callback:    commandMapB,
	},
	"explore": {
		name:        "explore",
		description: "shows you the pokemon in a given area",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "try to catch a pokemon",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "inspect the dex entry for the given pokemon",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "lists the names of all caught pokemons",
		callback:    commandPokedex,
	},
}

func startRepl(cfg *config) {
	cfg.supportedCommands = cliCommands
	cfg.PokeCache.NewCache(5 * time.Minute)
	cfg.PokeDex = make(map[string]Pokemon)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		input := cleanInput(scanner.Text())

		var commandName string
		var commandArg string
		if len(input) == 0 {
			commandName = ""
			commandArg = ""
		} else if len(input) == 1 {
			commandName = input[0]
			commandArg = ""
		} else {
			commandName = input[0]
			commandArg = input[1]
		}

		command, ok := cfg.supportedCommands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg, commandArg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	result := make([]string, 0)

	words := strings.SplitSeq(text, " ")
	for word := range words {
		if word == "" {
			continue
		}
		word = strings.ToLower(word)
		result = append(result, word)
	}

	return result
}
