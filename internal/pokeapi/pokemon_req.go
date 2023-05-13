package pokeapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + pokemonName
	var pokemon Pokemon
	if data, exists := c.cache.Get(url); exists {
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return pokemon, err
		}
		return pokemon, nil
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return pokemon, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return pokemon, err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return pokemon, nil
	}
	err = json.Unmarshal(bytes, &pokemon)
	if err != nil {
		return pokemon, nil
	}
	c.cache.Add(url, bytes)
	return pokemon, nil
}
