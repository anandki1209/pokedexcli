package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	// Default base url
	url := baseURL + "/location-area"
	// getting url for pagination
	if pageURL != nil {
		url = *pageURL
	}

	if data, ok := c.cache.Get(url); ok {
		var locationResp RespShallowLocations
		if err := json.Unmarshal(data, &locationResp); err != nil {
			return RespShallowLocations{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	if res.StatusCode == http.StatusOK {
		c.cache.Add(url, data)
	}

	var locationResp RespShallowLocations
	if err := json.Unmarshal(data, &locationResp); err != nil {
		return RespShallowLocations{}, err
	}

	return locationResp, nil

}
