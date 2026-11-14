package main

import (
	"encoding/json"
	"fmt"
)

func commandMapB(cfg *config, arg string) error {
	if cfg.Previous == "" {
		cfg.Previous = locationAreaEndpoint + "?offset=0&limit=20"
	}

	locations := &locationList{}

	val, ok := cfg.PokeCache.Get(cfg.Previous)
	var err error
	if !ok {
		err = getLocationAreas(cfg, locations, cfg.Previous)
		//fmt.Println("fetching")
	} else {
		err = json.Unmarshal(val, &locations)
		//fmt.Println("cache hit")
	}
	if err != nil {
		return err
	}

	cfg.Previous = locations.Previous
	cfg.Next = locations.Next

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
