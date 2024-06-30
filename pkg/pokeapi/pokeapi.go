package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Adam-Tay/pokecli/pkg/pokecache"
)

var cache pokecache.Cache = pokecache.NewCache(1 * time.Second)

func GetLocationsInPokemon(beginIndex, endIndex int) ([]PokemonLocation, error) {

	var locations []PokemonLocation
	// There's probably a better way to go about this
	if beginIndex < endIndex {
		for i := beginIndex; i <= endIndex; i++ {
			url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d", i)
			l, err := getLocationInPokemon(url)
			if err != nil {
				return locations, err
			}

			locations = append(locations, l)

		}
	} else {
		for i := beginIndex; i >= endIndex; i-- {
			url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d", i)
			l, err := getLocationInPokemon(url)
			if err != nil {
				return locations, err
			}

			locations = append(locations, l)
		}
	}
	return locations, nil
}

func GetPokemonByName(name string) (Pokemon, error) {
	p := Pokemon{}
	url := fmt.Sprint("https://pokeapi.co/api/v2/pokemon/" + name)

	cached, err := cache.Get(url)

	if err == nil {
		err := json.Unmarshal(cached.Value, &p)
		if err != nil {
			return p, err
		}
		return p, nil
	}

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
	cache.Add(url, []byte(body))
	return p, nil
}

func GetLocationInPokemonByName(name string) (PokemonLocation, error) {

	url := fmt.Sprint("https://pokeapi.co/api/v2/location-area/" + name)

	return getLocationInPokemon(url)
}

func getLocationInPokemon(url string) (PokemonLocation, error) {
	p := PokemonLocation{}

	cached, err := cache.Get(url)

	if err == nil {
		//var location PokemonLocation
		err := json.Unmarshal(cached.Value, &p)
		if err != nil {
			return p, err
		}
		return p, nil
	}

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
	cache.Add(url, []byte(body))
	return p, nil
}
