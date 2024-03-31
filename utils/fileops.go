package utils

import (
	"encoding/json"
	"os"
)

const JSON_FILE = "pokedex.json"

type JSONFileOps struct {
	pokedex map[string]PokemonProperties
}

type PokemonProperties struct {
	Name           string         `json:"name"`
	BaseExperience int            `json:"base_experience"`
	Height         int            `json:"height"`
	Weight         int            `json:"weight"`
	Stats          []PokemonStats `json:"stats"`
	Types          []PokemonType  `json:"types"`
}

type PokemonStats struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	}
}

type PokemonType struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

func (jr *JSONFileOps) Pokedex() map[string]PokemonProperties {
	return jr.pokedex
}

func (jr *JSONFileOps) readData(jsonFile string) (map[string]PokemonProperties, error) {
	file, err := os.Open(jsonFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var data map[string]PokemonProperties
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func (jr *JSONFileOps) writeData(jsonFile string, data map[string]PokemonProperties) error {
	file, err := os.Create(jsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func (jr *JSONFileOps) ReadPokedex() error {
	pokedex, err := jr.readData(JSON_FILE)
	if err != nil {
		return err
	}
	jr.pokedex = pokedex
	return nil
}

func (jr *JSONFileOps) WritePokedex(pokedex map[string]PokemonProperties) error {
	err := jr.writeData(JSON_FILE, pokedex)
	if err != nil {
		return err
	}
	return nil
}
