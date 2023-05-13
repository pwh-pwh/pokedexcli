package cli

import (
	"bufio"
	"fmt"
	"github.com/pwh-pwh/pokedexcli/config"
	"os"
	"strings"
)

func StartRepl(config *config.CliConfig) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(input string) []string {
	lower := strings.ToLower(input)
	return strings.Fields(lower)
}

type cliCommand struct {
	name        string
	description string
	callback    func(cliConfig *config.CliConfig) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCommand,
		},
		"map": {
			name:        "map",
			description: "List Some Location Areas",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous Location Areas",
			callback:    mapCommandBack,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitCommand,
		},
	}
}
