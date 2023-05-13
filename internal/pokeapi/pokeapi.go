package pokeapi

import (
	"github.com/pwh-pwh/pokedexcli/internal/pokecache"
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache      *pokecache.PokeCache
}

func NewClient(timeout time.Duration, cacheTime time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheTime),
	}
}
