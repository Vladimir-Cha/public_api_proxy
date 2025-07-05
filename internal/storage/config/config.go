package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type APIconfig struct {
	BaseURL    string `yaml:"base_url"`
	Timeout    time.Duration
	RawTimeout int `yaml:"timeout_seconds"`
}

type LoggingConfig struct {
	Enabled  bool   `yaml:"enabled"`
	LevelLog string `yaml:"level"`
}

type Config struct {
	API     APIconfig     `yaml:"api"`
	Logging LoggingConfig `yaml:"logging"`
}

// Загрузка конфига
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
