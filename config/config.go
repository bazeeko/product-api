package config

import (
	"encoding/json"
	"os"
)

type (
	Database struct {
		Driver   string `json:"driver"`
		FilePath string `json:"filepath"`
	}

	Config struct {
		Host     string   `json:"host"`
		Port     string   `json:"port"`
		Database Database `json:"database"`
	}
)

func NewConfig(configPath string) (config *Config, err error) {
	file, err := os.Open(configPath)
	if err != nil {
		return
	}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return
	}

	return
}
