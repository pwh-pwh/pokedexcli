package tests

import (
	"fmt"
	"github.com/pwh-pwh/pokedexcli/internal/pokeapi"
	"log"
	"testing"
)

func TestListLocationAreas(t *testing.T) {
	client := pokeapi.NewClient()
	areas, err := client.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(areas)
}
