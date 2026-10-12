package pokeapi

import (
	"encoding/json"
	"net/http"
)

type PokeResponse struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous any       `json:"previous"`
	Results  []Results `json:"results"`
}
type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocationAreas(url string) (PokeResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return PokeResponse{}, err
	}
	defer res.Body.Close()
	var responseResult PokeResponse

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&responseResult); err != nil {
		return PokeResponse{}, err
	}

	return responseResult, nil

}
