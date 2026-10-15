package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(name string) (Location, error) {
	url := baseURL + "/location-area/" + name
	if data, ok := c.cache.Get(url); ok {
		var location Location
		if err := json.Unmarshal(data, &location); err != nil {
			return Location{}, err
		}
		return location, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	if res.StatusCode == http.StatusOK {
		c.cache.Add(url, data)
	}

	var location Location

	if err := json.Unmarshal(data, &location); err != nil {
		return Location{}, err
	}

	return location, nil

}
