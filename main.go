package main

import (
	"github.com/pwh-pwh/pokedexcli/cli"
	"github.com/pwh-pwh/pokedexcli/config"
	"github.com/pwh-pwh/pokedexcli/internal/pokeapi"
)

func main() {
	config := config.CliConfig{
		PokeApiClient: pokeapi.NewClient(),
		PreviousUrl:   nil,
		NextUrl:       nil,
	}
	cli.StartRepl(&config)
}
