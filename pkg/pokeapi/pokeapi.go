package pokeapi

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io"
)

type PokemonLocation struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocationsInPokemon(beginIndex, endIndex int)  ([]PokemonLocation, error) {
	var locations []PokemonLocation 
	// There's probably a better way to go about this
	if beginIndex < endIndex {
		for i := beginIndex; i <= endIndex; i++ {
			l, err := getLocationInPokemon(i)
			if err != nil {
				return locations, err
			}

			locations = append(locations, l)

		}
	} else {
		for i := beginIndex; i >= endIndex; i-- {
			l, err := getLocationInPokemon(i)
			if err != nil {
				return locations, err
			}

			locations = append(locations, l)
		}
	}
	return locations, nil
}

func getLocationInPokemon(locationId int) (PokemonLocation, error) {
	p := PokemonLocation{}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d", locationId)
	resp, err := http.Get(url)
	if err != nil {
		return p, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return p, err
	}

	
	err = json.Unmarshal(body, &p)
	if err != nil {
		return p, err
	}

	return p, nil
}
/*
Chat-gpt Prompt:
In a golang project I have a main.go file, then I have a pkg/pokeapi/pokeapi.go file which is going to contain functions for calling the PokeAPI. How should I set this pokeapi.go file up?

*/