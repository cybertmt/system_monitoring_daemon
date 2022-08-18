package config

import (
	"fmt"
	"os"

	yamlv3 "gopkg.in/yaml.v3"
)

type Config struct {
	Stats StatsConfig
}

type StatsConfig struct {
	LoadAvg bool
	CPU     bool
	Disk    bool
	NetTop  bool
	NetStat bool
}

func NewConfig() Config {
	return Config{}
}

func LoadConfig(path string) (*Config, error) {
	configContent, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("invalid config file %s: %w", path, err)
	}

	config := NewConfig()
	err = yamlv3.Unmarshal(configContent, &config)
	if err != nil {
		return nil, fmt.Errorf("invalid config file content %s: %w", path, err)
	}

	return &config, nil
}
