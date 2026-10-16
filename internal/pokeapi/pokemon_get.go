package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {

	url := baseURL + "/pokemon/" + name

	if data, ok := c.cache.Get(url); ok {
		var pokemon Pokemon
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}

	if res.StatusCode == http.StatusNotFound {
		return Pokemon{}, errors.New("No pokemon exists with this name")
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	if res.StatusCode == http.StatusOK {
		c.cache.Add(url, data)
	}

	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil

}
