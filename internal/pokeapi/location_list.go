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

	if data, ok := c.httpCache.Get(url); ok {
		return unmarshalData(data)
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
		c.httpCache.Add(url, data)
	}

	return unmarshalData(data)

}

func unmarshalData(data []byte) (RespShallowLocations, error) {
	var locationResp RespShallowLocations
	if err := json.Unmarshal(data, &locationResp); err != nil {
		return RespShallowLocations{}, err
	}

	return locationResp, nil
}
