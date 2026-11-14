package main

import (
	"fmt"
)

func commandPokedex(cfg *config, arg string) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range cfg.PokeDex {
		fmt.Printf(" - %s\n", key)
	}
	return nil
}
