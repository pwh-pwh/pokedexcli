package cli

import (
	"errors"
	"fmt"
	"github.com/pwh-pwh/pokedexcli/config"
	"math/rand"
)

func catchCommand(config *config.CliConfig, args ...string) error {
	if len(args) != 1 {
		return errors.New("must input pokemon name")
	}
	pokemon, err := config.PokeApiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", args[0])
	rNum := rand.Intn(pokemon.BaseExperience)
	if rNum < 40 {
		fmt.Printf("%v was caught!", args[0])
		config.CaughtPokemon[args[0]] = pokemon
	} else {
		fmt.Printf("%v escaped!", args[0])
	}
	fmt.Println()
	return err
}
