package cli

import (
	"fmt"
	"github.com/pwh-pwh/pokedexcli/config"
)

func pokedexCommand(config *config.CliConfig, args ...string) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range config.CaughtPokemon {
		fmt.Printf(" - %v\n", key)
	}
	return nil
}
