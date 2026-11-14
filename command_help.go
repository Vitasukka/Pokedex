package main

import "fmt"

func commandHelp(cfg *config, arg string) error {
	message := fmt.Sprintf("Welcome to the Pokedex!\nUsage:\n\n")

	for _, val := range cfg.supportedCommands {
		message += fmt.Sprintf("%s: %s\n", val.name, val.description)
	}

	fmt.Print(message)
	return nil
}
