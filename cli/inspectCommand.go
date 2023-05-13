package cli

import (
	"errors"
	"fmt"
	"github.com/pwh-pwh/pokedexcli/config"
)

func inspectCommand(config *config.CliConfig, args ...string) error {
	if len(args) != 1 {
		return errors.New("please input pokemon name")
	}
	if pokemon, exists := config.CaughtPokemon[args[0]]; exists {
		fmt.Printf("Name:%v\nHeight:%v\nWeight:%v\n", pokemon.Name, pokemon.Height, pokemon.Weight)
		fmt.Println("Stats:")
		for _, s := range pokemon.Stats {
			fmt.Printf("  -%v: %v\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("  - %v\n", t.Type.Name)
		}
	} else {
		fmt.Print("you have not caught that pokemon\n")
	}
	return nil
}
