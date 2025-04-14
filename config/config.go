package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Env struct {
		Port      string `yaml:"PORT"`
		Model     string `yaml:"MODEL"`
		OllamaUrl string `yaml:"OLLAMAURL"`
		SearchUrl string `yaml:"SEARCHURL"`
	} `yaml:"env"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config

	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
