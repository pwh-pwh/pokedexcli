package cli

import (
	"errors"
	"fmt"
	"github.com/pwh-pwh/pokedexcli/config"
)

func mapCommand(config *config.CliConfig, args ...string) error {
	resp, err := config.PokeApiClient.ListLocationAreas(config.NextUrl)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, location := range resp.Results {
		fmt.Printf(" - %s\n", location.Name)
	}
	config.NextUrl = resp.Next
	config.PreviousUrl = resp.Previous
	return nil
}

func mapCommandBack(config *config.CliConfig, _ ...string) error {
	if config.PreviousUrl == nil {
		return errors.New("you are in first page")
	}
	resp, err := config.PokeApiClient.ListLocationAreas(config.PreviousUrl)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, location := range resp.Results {
		fmt.Printf(" - %s\n", location.Name)
	}
	config.NextUrl = resp.Next
	config.PreviousUrl = resp.Previous
	return nil
}
