package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/siddhant-vij/Pokedex/api"
)

func CatchPokemon(pokemon string) {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	err := catchPokemon(pokemon)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println()
		return
	}

	fmt.Println()
}

func catchPokemon(pokemon string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)

	body, err := api.FetchDataFromAPI(url)
	if err != nil {
		return err
	}

	var response PokemonProperties

	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	baseExperience := response.BaseExperience
	randomDifficult := rand.Intn(baseExperience)

	if randomDifficult > 40 {
		fmt.Printf("%s was caught!\n", response.Name)
		pokedex[response.Name] = response
	} else {
		fmt.Printf("%s escaped.\n", response.Name)
	}

	return nil
}
