package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/chzyer/readline"

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

	rl, err := readline.New("pokedex > ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}

		line = strings.ToLower(strings.TrimSpace(line))
		if line == "" {
			continue
		}

		rl.SaveHistory(line)

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
			fmt.Println()
		case "clear":
			fmt.Print("\033[2J\033[1;1H")
		default:
			fmt.Printf("Unknown command: %s\n\n", line)
		}
	}
}
