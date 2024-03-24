package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/siddhant-vij/Pokedex/api"
	"github.com/siddhant-vij/Pokedex/config"
)

func Run() {
	cfg := &config.Config{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		if scanner.Scan() {
			line := strings.ToLower(strings.TrimSpace(scanner.Text()))
			switch line {
			case "map":
				displayLocationAreas(cfg, true)
			case "mapb":
				displayLocationAreas(cfg, false)
			case "exit":
				os.Exit(0)
			case "help":
				printHelp()
			default:
				fmt.Printf("Unknown command: %s\n", line)
			}
		}
	}
}

func displayLocationAreas(cfg *config.Config, isNext bool) {
	areas, err := api.FetchLocationAreas(cfg, isNext)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println()
		return
	}

	for _, area := range areas {
		fmt.Println(area.Name)
	}
	fmt.Println()
}

func printHelp() {
	fmt.Println("Commands:")
	fmt.Println("  help: Displays this help message")
	fmt.Println("  exit: Exits the Pokedex")
	fmt.Println("  map: Displays the names of next 20 location areas in the Pokemon world")
	fmt.Println("  mapb: Displays the names of previous 20 location areas in the Pokemon world")
}
