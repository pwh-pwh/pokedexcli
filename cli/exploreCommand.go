package cli

import (
	"errors"
	"fmt"
	"github.com/pwh-pwh/pokedexcli/config"
)

func exploreCommand(config *config.CliConfig, args ...string) error {
	if len(args) != 1 {
		return errors.New("args len err")
	}
	location, err := config.PokeApiClient.LocationGet(args[0])
	if err != nil {
		return err
	}
	fmt.Println("Exploring pastoria-city-area...")
	fmt.Println("Found Pokemon: ")
	for _, po := range location.PokemonEncounters {
		fmt.Printf(" - %v\n", po.Pokemon.Name)
	}
	return nil
}
