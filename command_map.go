package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/vitasukka/pokedex/pokecache"
)

type locationList struct {
	Count    int    `json:count`
	Next     string `json:next`
	Previous string `json:previous`
	Results  []struct {
		Name string `json:name`
		Url  string `json:url`
	} `json:results`
}

const locationAreaEndpoint = "https://pokeapi.co/api/v2/location-area/"

func commandMap(cfg *config, arg string) error {
	if cfg.Next == "" {
		cfg.Next = locationAreaEndpoint + "?offset=0&limit=20"
	}

	locations := &locationList{}

	val, ok := cfg.PokeCache.Get(cfg.Next)
	var err error
	if !ok {
		err = getLocationAreas(cfg, locations, cfg.Next)
	} else {
		err = json.Unmarshal(val, &locations)
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

func getLocationAreas(cfg *config, locations *locationList, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	cfg.PokeCache.Add(url, pokecache.CacheEntry{Val: data})

	err = json.Unmarshal(data, &locations)
	if err != nil {
		return err
	}

	return nil
}
