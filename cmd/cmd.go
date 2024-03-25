package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/siddhant-vij/Pokedex/config"
)

var pokedex = map[string]PokemonProperties{}

type PokemonProperties struct {
	Name           string         `json:"name"`
	BaseExperience int            `json:"base_experience"`
	Height         int            `json:"height"`
	Weight         int            `json:"weight"`
	Stats          []PokemonStats `json:"stats"`
	Types          []PokemonType  `json:"types"`
}

type PokemonStats struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	}
}

type PokemonType struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

func Run() {
	cfg := &config.Config{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		if scanner.Scan() {
			line := strings.ToLower(strings.TrimSpace(scanner.Text()))
			args := strings.Split(line, " ")
			switch args[0] {
			case "map":
				DisplayLocationAreas(cfg, true)
			case "mapb":
				DisplayLocationAreas(cfg, false)
			case "explore":
				if len(args) < 2 {
					fmt.Println("Usage: explore <location>")
					fmt.Println()
					continue
				}
				ExplorePokemons(args[1])
			case "catch":
				if len(args) < 2 {
					fmt.Println("Usage: catch <pokemon>")
					fmt.Println()
					continue
				}
				CatchPokemon(args[1])
			case "inspect":
				if len(args) < 2 {
					fmt.Println("Usage: inspect <pokemon>")
					fmt.Println()
					continue
				}
				InspectPokemon(args[1])
			case "pokedex":
				fmt.Println("Your Pokedex:")
				if len(pokedex) == 0 {
					fmt.Println("  <empty>")
				}
				for name := range pokedex {
					fmt.Println("  -", name)
				}
				fmt.Println()
			case "exit":
				os.Exit(0)
			case "help":
				PrintHelp()
			case "clear":
				fmt.Print("\033[2J\033[1;1H")
			default:
				fmt.Printf("Unknown command: %s\n", line)
			}
		}
	}
}
