package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/siddhant-vij/Pokedex/api"
)

var pokemonCache = api.NewCache(1*time.Minute, 5*time.Minute)

type Pokemon struct {
	Pokemon PokemonDetails `json:"pokemon"`
}

type PokemonDetails struct {
	Name string `json:"name"`
}

func ExplorePokemons(location string) {
	pokemons, err := fetchPokemons(location)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println()
		return
	}

	for _, pokemon := range pokemons {
		fmt.Println(pokemon.Pokemon.Name)
	}
	fmt.Println()
}

func fetchPokemons(location string) ([]Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)

	var cachedResponse []Pokemon

	if data, _, _, found := pokemonCache.Get(url); found {
		fmt.Println("Serving from cache...")
		if err := json.Unmarshal(data, &cachedResponse); err == nil {
			return cachedResponse, nil
		}
		fmt.Println("Cache found but unable to unmarshal, fetching from API...")
	}

	body, err := api.FetchDataFromAPI(url)
	if err != nil {
		return nil, err
	}

	var response struct {
		Results []Pokemon `json:"pokemon_encounters"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	responseBody, err := json.Marshal(response.Results)
	if err != nil {
		return nil, fmt.Errorf("error marshaling response: %w", err)
	}

	pokemonCache.Add(url, responseBody, "", "")

	return response.Results, nil
}
