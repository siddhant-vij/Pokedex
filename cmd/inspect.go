package cmd

import (
	"fmt"

	"github.com/siddhant-vij/Pokedex/api"
)

func InspectPokemon(pokemon string) {
	if _, found := pokedex[pokemon]; !found {
		url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)
		_, err := api.FetchDataFromAPI(url)
		if err != nil {
			fmt.Printf("Check if %s exists as a Pokemon...\n", pokemon)
		} else {
			fmt.Printf("You've not caught %s yet...\n", pokemon)
		}
	} else {
		printPokemonDetails(pokemon)
	}

	fmt.Println()
}

func printPokemonDetails(pokemon string) {
	fmt.Printf("Name: %s\n", pokedex[pokemon].Name)
	fmt.Printf("Base experience: %d\n", pokedex[pokemon].BaseExperience)
	fmt.Printf("Height: %d\n", pokedex[pokemon].Height)
	fmt.Printf("Weight: %d\n", pokedex[pokemon].Weight)
	fmt.Println("Stats:")
	for _, stat := range pokedex[pokemon].Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, type_ := range pokedex[pokemon].Types {
		fmt.Printf("  -%s\n", type_.Type.Name)
	}
}
