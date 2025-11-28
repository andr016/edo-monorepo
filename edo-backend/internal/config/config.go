package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	TmpPath string `json:"tmpPath"`
}

func LoadConfig() (Config, error) {
	var config Config
	data, err := os.ReadFile("config.json")
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	return config, err
}
