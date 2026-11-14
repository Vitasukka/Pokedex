package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/vitasukka/pokedex/pokecache"
)

type LocationArea struct {
	Id                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel int `json:"min_level"`
				MaxLevel int `json:"max_level"`
				Chance   int `json:"chance"`
				Method   struct {
					Name string `json:"name"`
					Url  string `json:"url"`
				} `json:"method"`
				ConditionValues []interface{} `json:"condition_values"`
				PokemonDetails  interface{}   `json:"pokemon_details"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cfg *config, arg string) error {
	url := locationAreaEndpoint + "/" + arg + "/"

	location := &LocationArea{}
	val, ok := cfg.PokeCache.Get(url)
	var err error
	if !ok {
		err = getLocationArea(cfg, location, url)
	} else {
		err = json.Unmarshal(val, location)
	}
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\nFound Pokemon:\n", arg)
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func getLocationArea(cfg *config, location *LocationArea, url string) error {
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

	err = json.Unmarshal(data, &location)
	if err != nil {
		return err
	}

	return nil
}
