package main

import (
	"github.com/pwh-pwh/pokedexcli/cli"
	"github.com/pwh-pwh/pokedexcli/config"
	"github.com/pwh-pwh/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	config := config.CliConfig{
		PokeApiClient: pokeapi.NewClient(time.Minute, time.Second),
		PreviousUrl:   nil,
		NextUrl:       nil,
		CaughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	cli.StartRepl(&config)
}
