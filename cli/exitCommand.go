package cli

import (
	"github.com/pwh-pwh/pokedexcli/config"
	"os"
)

func exitCommand(config *config.CliConfig, args ...string) error {
	os.Exit(0)
	return nil
}
