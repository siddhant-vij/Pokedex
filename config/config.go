package config

type Config struct {
	Next string `json:"next,omitempty"`
	Prev string `json:"previous,omitempty"`
	Current string `json:"current,omitempty"`
}
