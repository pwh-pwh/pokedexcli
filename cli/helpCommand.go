package cli

import (
	"fmt"
	"github.com/pwh-pwh/pokedexcli/config"
)

func helpCommand(config *config.CliConfig) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
