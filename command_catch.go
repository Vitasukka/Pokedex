package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

type Pokemon struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	PastAbilities []struct {
		Generation struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"generation"`
		Abilities []struct {
			IsHidden bool        `json:"is_hidden"`
			Slot     int         `json:"slot"`
			Ability  interface{} `json:"ability"`
		} `json:"abilities"`
	} `json:"past_abilities"`
	Forms []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	HeldItems              []interface{} `json:"held_items"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt int `json:"level_learned_at"`
			VersionGroup   struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version_group"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"move_learn_method"`
			Order *int `json:"order"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Species struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		Other struct {
			Home struct {
				FrontShiny       string      `json:"front_shiny"`
				FrontFemale      interface{} `json:"front_female"`
				FrontDefault     string      `json:"front_default"`
				FrontShinyFemale interface{} `json:"front_shiny_female"`
			} `json:"home"`
			Showdown struct {
				BackShiny        string      `json:"back_shiny"`
				BackFemale       interface{} `json:"back_female"`
				FrontShiny       string      `json:"front_shiny"`
				BackDefault      string      `json:"back_default"`
				FrontFemale      interface{} `json:"front_female"`
				FrontDefault     string      `json:"front_default"`
				BackShinyFemale  interface{} `json:"back_shiny_female"`
				FrontShinyFemale interface{} `json:"front_shiny_female"`
			} `json:"showdown"`
			DreamWorld struct {
				FrontFemale  interface{} `json:"front_female"`
				FrontDefault string      `json:"front_default"`
			} `json:"dream_world"`
			OfficialArtwork struct {
				Versions struct {
					GenerationI struct {
						RedAndBlue struct {
							FrontDefault interface{} `json:"front_default"`
						} `json:"red-and-blue"`
						RedAndGreen struct {
							FrontDefault interface{} `json:"front_default"`
						} `json:"red-and-green"`
					} `json:"generation-i"`
					GenerationIi struct {
						GoldAndSilver struct {
							FrontDefault string `json:"front_default"`
						} `json:"gold-and-silver"`
					} `json:"generation-ii"`
				} `json:"versions"`
				FrontShiny   string `json:"front_shiny"`
				FrontDefault string `json:"front_default"`
			} `json:"official-artwork"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				Yellow struct {
					BackGbc              interface{} `json:"back_gbc"`
					BackGray             interface{} `json:"back_gray"`
					FrontGbc             interface{} `json:"front_gbc"`
					FrontGray            interface{} `json:"front_gray"`
					BackDefault          interface{} `json:"back_default"`
					FrontDefault         interface{} `json:"front_default"`
					BackTransparent      interface{} `json:"back_transparent"`
					FrontTransparent     interface{} `json:"front_transparent"`
					BackTransparentGray  interface{} `json:"back_transparent_gray"`
					FrontTransparentGray interface{} `json:"front_transparent_gray"`
				} `json:"yellow"`
				RedBlue struct {
					BackGray             interface{} `json:"back_gray"`
					FrontGray            interface{} `json:"front_gray"`
					BackDefault          interface{} `json:"back_default"`
					FrontDefault         interface{} `json:"front_default"`
					BackTransparent      interface{} `json:"back_transparent"`
					FrontTransparent     interface{} `json:"front_transparent"`
					BackTransparentGray  interface{} `json:"back_transparent_gray"`
					FrontTransparentGray interface{} `json:"front_transparent_gray"`
				} `json:"red-blue"`
				RedGreenJapan struct {
					BackGray     interface{} `json:"back_gray"`
					FrontGray    interface{} `json:"front_gray"`
					BackDefault  interface{} `json:"back_default"`
					FrontDefault interface{} `json:"front_default"`
				} `json:"red-green-japan"`
			} `json:"generation-i"`
			GenerationV struct {
				Icons struct {
					Animated struct {
						FrontDefault string `json:"front_default"`
					} `json:"animated"`
					FrontDefault string `json:"front_default"`
				} `json:"icons"`
				BlackWhite struct {
					Animated struct {
						BackShiny        string      `json:"back_shiny"`
						BackFemale       interface{} `json:"back_female"`
						FrontShiny       string      `json:"front_shiny"`
						BackDefault      string      `json:"back_default"`
						FrontFemale      interface{} `json:"front_female"`
						FrontDefault     string      `json:"front_default"`
						BackShinyFemale  interface{} `json:"back_shiny_female"`
						FrontShinyFemale interface{} `json:"front_shiny_female"`
					} `json:"animated"`
					BackShiny        string      `json:"back_shiny"`
					BackFemale       interface{} `json:"back_female"`
					FrontShiny       string      `json:"front_shiny"`
					BackDefault      string      `json:"back_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     string      `json:"front_default"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationIi struct {
				Gold struct {
					BackShiny             string `json:"back_shiny"`
					FrontShiny            string `json:"front_shiny"`
					BackDefault           string `json:"back_default"`
					FrontDefault          string `json:"front_default"`
					BackTransparent       string `json:"back_transparent"`
					FrontTransparent      string `json:"front_transparent"`
					BackShinyTransparent  string `json:"back_shiny_transparent"`
					FrontShinyTransparent string `json:"front_shiny_transparent"`
				} `json:"gold"`
				Silver struct {
					BackShiny             string `json:"back_shiny"`
					FrontShiny            string `json:"front_shiny"`
					BackDefault           string `json:"back_default"`
					FrontDefault          string `json:"front_default"`
					BackTransparent       string `json:"back_transparent"`
					FrontTransparent      string `json:"front_transparent"`
					BackShinyTransparent  string `json:"back_shiny_transparent"`
					FrontShinyTransparent string `json:"front_shiny_transparent"`
				} `json:"silver"`
				Crystal struct {
					Animated struct {
						FrontShiny   string `json:"front_shiny"`
						FrontDefault string `json:"front_default"`
					} `json:"animated"`
					BackShiny             string `json:"back_shiny"`
					FrontShiny            string `json:"front_shiny"`
					BackDefault           string `json:"back_default"`
					FrontDefault          string `json:"front_default"`
					BackTransparent       string `json:"back_transparent"`
					FrontTransparent      string `json:"front_transparent"`
					BackShinyTransparent  string `json:"back_shiny_transparent"`
					FrontShinyTransparent string `json:"front_shiny_transparent"`
				} `json:"crystal"`
			} `json:"generation-ii"`
			GenerationIv struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
				} `json:"icons"`
				Platinum struct {
					Animated struct {
						FrontShiny       string      `json:"front_shiny"`
						FrontFemale      interface{} `json:"front_female"`
						FrontDefault     string      `json:"front_default"`
						FrontShinyFemale interface{} `json:"front_shiny_female"`
					} `json:"animated"`
					BackShiny        string      `json:"back_shiny"`
					BackFemale       interface{} `json:"back_female"`
					FrontShiny       string      `json:"front_shiny"`
					BackDefault      string      `json:"back_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     string      `json:"front_default"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"platinum"`
				DiamondPearl struct {
					Animated struct {
						FrontShiny       string      `json:"front_shiny"`
						FrontFemale      interface{} `json:"front_female"`
						FrontDefault     string      `json:"front_default"`
						FrontShinyFemale interface{} `json:"front_shiny_female"`
					} `json:"animated"`
					BackShiny        string      `json:"back_shiny"`
					BackFemale       interface{} `json:"back_female"`
					FrontShiny       string      `json:"front_shiny"`
					BackDefault      string      `json:"back_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     string      `json:"front_default"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					Animated struct {
						FrontShiny       string      `json:"front_shiny"`
						FrontFemale      interface{} `json:"front_female"`
						FrontDefault     string      `json:"front_default"`
						FrontShinyFemale interface{} `json:"front_shiny_female"`
					} `json:"animated"`
					BackShiny        string      `json:"back_shiny"`
					BackFemale       interface{} `json:"back_female"`
					FrontShiny       string      `json:"front_shiny"`
					BackDefault      string      `json:"back_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     string      `json:"front_default"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
			} `json:"generation-iv"`
			GenerationIx struct {
				Champions struct {
					Icons struct {
						FrontShiny   interface{} `json:"front_shiny"`
						FrontDefault interface{} `json:"front_default"`
					} `json:"icons"`
					BackShiny        interface{} `json:"back_shiny"`
					BackFemale       interface{} `json:"back_female"`
					FrontShiny       interface{} `json:"front_shiny"`
					BackDefault      interface{} `json:"back_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     interface{} `json:"front_default"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"champions"`
				LegendsZa struct {
					Icons struct {
						FrontShiny       string      `json:"front_shiny"`
						FrontFemale      interface{} `json:"front_female"`
						FrontDefault     string      `json:"front_default"`
						FrontShinyFemale interface{} `json:"front_shiny_female"`
					} `json:"icons"`
					BackShiny        string      `json:"back_shiny"`
					BackFemale       interface{} `json:"back_female"`
					FrontShiny       string      `json:"front_shiny"`
					BackDefault      string      `json:"back_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     string      `json:"front_default"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"legends-za"`
				ScarletViolet struct {
					Icons struct {
						FrontFemale  interface{} `json:"front_female"`
						FrontDefault string      `json:"front_default"`
					} `json:"icons"`
					BackShiny        string      `json:"back_shiny"`
					BackFemale       interface{} `json:"back_female"`
					FrontShiny       string      `json:"front_shiny"`
					BackDefault      string      `json:"back_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     string      `json:"front_default"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"scarlet-violet"`
			} `json:"generation-ix"`
			GenerationVi struct {
				XY struct {
					BackShiny        string      `json:"back_shiny"`
					BackFemale       interface{} `json:"back_female"`
					FrontShiny       string      `json:"front_shiny"`
					BackDefault      string      `json:"back_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     string      `json:"front_default"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"x-y"`
				Icons struct {
					FrontFemale  interface{} `json:"front_female"`
					FrontDefault string      `json:"front_default"`
				} `json:"icons"`
				OmegaRubyAlphaSapphire struct {
					BackShiny        string      `json:"back_shiny"`
					BackFemale       interface{} `json:"back_female"`
					FrontShiny       string      `json:"front_shiny"`
					BackDefault      string      `json:"back_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     string      `json:"front_default"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"omega-ruby-alpha-sapphire"`
			} `json:"generation-vi"`
			GenerationIii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
				} `json:"icons"`
				Emerald struct {
					Animated struct {
						BackShiny    string `json:"back_shiny"`
						FrontShiny   string `json:"front_shiny"`
						BackDefault  string `json:"back_default"`
						FrontDefault string `json:"front_default"`
					} `json:"animated"`
					BackShiny    string `json:"back_shiny"`
					FrontShiny   string `json:"front_shiny"`
					BackDefault  string `json:"back_default"`
					FrontDefault string `json:"front_default"`
				} `json:"emerald"`
				RubySapphire struct {
					BackShiny    string `json:"back_shiny"`
					FrontShiny   string `json:"front_shiny"`
					BackDefault  string `json:"back_default"`
					FrontDefault string `json:"front_default"`
				} `json:"ruby-sapphire"`
				FireredLeafgreen struct {
					BackShiny    string `json:"back_shiny"`
					FrontShiny   string `json:"front_shiny"`
					BackDefault  string `json:"back_default"`
					FrontDefault string `json:"front_default"`
				} `json:"firered-leafgreen"`
			} `json:"generation-iii"`
			GenerationVii struct {
				Icons struct {
					FrontFemale  interface{} `json:"front_female"`
					FrontDefault string      `json:"front_default"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					BackShiny        string      `json:"back_shiny"`
					BackFemale       interface{} `json:"back_female"`
					FrontShiny       string      `json:"front_shiny"`
					BackDefault      string      `json:"back_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     string      `json:"front_default"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
				LetsGoPikachuLetsGoEevee struct {
					Icons struct {
						FrontDefault interface{} `json:"front_default"`
					} `json:"icons"`
					BackShiny        interface{} `json:"back_shiny"`
					BackFemale       interface{} `json:"back_female"`
					FrontShiny       interface{} `json:"front_shiny"`
					BackDefault      interface{} `json:"back_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     interface{} `json:"front_default"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"lets-go-pikachu-lets-go-eevee"`
			} `json:"generation-vii"`
			GenerationViii struct {
				SwordShield struct {
					Icons struct {
						FrontFemale  interface{} `json:"front_female"`
						FrontDefault string      `json:"front_default"`
					} `json:"icons"`
					BackShiny        string      `json:"back_shiny"`
					BackFemale       interface{} `json:"back_female"`
					FrontShiny       string      `json:"front_shiny"`
					BackDefault      string      `json:"back_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     string      `json:"front_default"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"sword-shield"`
				LegendsArceus struct {
					Icons struct {
						FrontShiny       interface{} `json:"front_shiny"`
						FrontFemale      interface{} `json:"front_female"`
						FrontDefault     interface{} `json:"front_default"`
						FrontShinyFemale interface{} `json:"front_shiny_female"`
					} `json:"icons"`
					BackShiny        interface{} `json:"back_shiny"`
					BackFemale       interface{} `json:"back_female"`
					FrontShiny       interface{} `json:"front_shiny"`
					BackDefault      interface{} `json:"back_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     interface{} `json:"front_default"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"legends-arceus"`
				BrilliantDiamondShiningPearl struct {
					Icons struct {
						FrontDefault string `json:"front_default"`
					} `json:"icons"`
					BackShiny        string      `json:"back_shiny"`
					BackFemale       interface{} `json:"back_female"`
					FrontShiny       string      `json:"front_shiny"`
					BackDefault      string      `json:"back_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     string      `json:"front_default"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"brilliant-diamond-shining-pearl"`
			} `json:"generation-viii"`
		} `json:"versions"`
		BackShiny        string      `json:"back_shiny"`
		BackFemale       interface{} `json:"back_female"`
		FrontShiny       string      `json:"front_shiny"`
		BackDefault      string      `json:"back_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontDefault     string      `json:"front_default"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"sprites"`
	Cries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	PastStats []interface{} `json:"past_stats"`
	Types     []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	PastTypes []interface{} `json:"past_types"`
}

const pokemonEndpoint = "https://pokeapi.co/api/v2/pokemon/"

func commandCatch(cfg *config, arg string) error {
	arg = strings.ToLower(arg)
	url := pokemonEndpoint + arg + "/"

	pokemon := &Pokemon{}
	err := getPokemon(cfg, pokemon, url)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", arg)
	if rand.Intn(255) >= pokemon.BaseExperience {
		cfg.PokeDex[arg] = *pokemon
		fmt.Printf("%s was caught!\n", arg)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped\n", arg)
	}

	return nil
}

func getPokemon(cfg *config, p *Pokemon, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(p)
	if err != nil {
		return err
	}

	return nil
}
