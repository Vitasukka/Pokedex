package main

import (
	"fmt"
	"strings"
)

func commandInspect(cfg *config, arg string) error {
	arg = strings.ToLower(arg)
	pokemon, ok := cfg.PokeDex[arg]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", arg)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}

    return nil
}
