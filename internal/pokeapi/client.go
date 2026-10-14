package pokeapi

import (
	"net/http"
	"time"

	"github.com/anandki1209/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	httpCache  *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		httpCache: pokecache.NewCache(5 * time.Second),
	}
}
