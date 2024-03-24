package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/siddhant-vij/Pokedex/config"
)

var pokeCache = NewCache(1*time.Minute, 5*time.Minute)

type LocationArea struct {
	Name string `json:"name"`
}

func FetchLocationAreas(cfg *config.Config, isNext bool) ([]LocationArea, error) {
	url := determineURL(cfg, isNext)
	if url == "" {
		return nil, fmt.Errorf("no further data")
	}

	var cachedResponse []LocationArea

	if data, next, prev, found := pokeCache.Get(url); found {
		fmt.Println("Serving from cache...")
		if err := json.Unmarshal(data, &cachedResponse); err == nil {
			updateConfigURLs(cfg, next, prev, isNext)
			return cachedResponse, nil
		}
		fmt.Println("Cache found but unable to unmarshal, fetching from API...")
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var response struct {
		Results []LocationArea `json:"results"`
		Next    string         `json:"next,omitempty"`
		Prev    string         `json:"previous,omitempty"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	updateConfigURLs(cfg, response.Next, response.Prev, isNext)

	responseBody, err := json.Marshal(response.Results)
	if err != nil {
		return nil, fmt.Errorf("error marshaling response: %w", err)
	}
	pokeCache.Add(url, responseBody, response.Next, response.Prev)
	
	return response.Results, nil
}

func updateConfigURLs(cfg *config.Config, nextURL, prevURL string, isNext bool) {
	if isNext {
		cfg.Prev = cfg.Current
		cfg.Current = nextURL
		cfg.Next = nextURL
	} else {
		cfg.Next = cfg.Current
		cfg.Current = prevURL
		cfg.Prev = prevURL
	}
}

func determineURL(cfg *config.Config, isNext bool) string {
	if isNext {
		if cfg.Next == "" {
			return "https://pokeapi.co/api/v2/location-area?limit=20"
		}
		return cfg.Next
	} else {
		return cfg.Prev
	}
}
