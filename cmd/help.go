package cmd

import "fmt"

func PrintHelp() {
	fmt.Println("Commands:")
	fmt.Println("  map: Displays the names of next 20 location areas in the Pokemon world")
	fmt.Println("  mapb: Displays the names of previous 20 location areas in the Pokemon world")
	fmt.Println("  explore <location>: Displays the pokemons in the specified location")
	fmt.Println("  catch <pokemon>: Tries to catch the specified pokemon")
	fmt.Println("  inspect <pokemon>: Displays the details of the specified pokemon if caught")
	fmt.Println("  pokedex: Displays your pokedex")
	fmt.Println("  exit: Exits the Pokedex")
	fmt.Println("  help: Displays this help message")
	fmt.Println("  clear: Clears the screen")
}
