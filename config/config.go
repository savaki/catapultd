package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	AgentId   string `json:"agent-id"`
	AuthToken string `json:"auth-token"`
}

func Load(r io.Reader) (*Config, error) {
	config := &Config{}
	err := json.NewDecoder(r).Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func LoadFile(filename string) (*Config, error) {
	r, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return Load(r)
}
