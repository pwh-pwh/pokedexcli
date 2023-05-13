package config

import "github.com/pwh-pwh/pokedexcli/internal/pokeapi"

type CliConfig struct {
	PokeApiClient pokeapi.Client
	PreviousUrl   *string
	NextUrl       *string
}
