package tests

import (
	"fmt"
	"github.com/pwh-pwh/pokedexcli/internal/pokeapi"
	"log"
	"testing"
	"time"
)

func TestGetPokemon(t *testing.T) {
	client := pokeapi.NewClient(time.Minute, time.Minute)
	data, err := client.GetPokemon("bulbasaur")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
