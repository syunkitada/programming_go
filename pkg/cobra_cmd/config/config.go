package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Api   ApiConfig
	Agent AgentConfig
}

type ApiConfig struct {
	Listen   string
	Interval int
}

type AgentConfig struct {
	ReportInterval int
}

func NewConfig() *Config {
	defaultConfig := &Config{
		Api: ApiConfig{
			Listen:   "0.0.0.0:8080",
			Interval: 10,
		},
		Agent: AgentConfig{
			ReportInterval: 10,
		},
	}
	return defaultConfig
}

var Conf Config

func LoadConfig(configFile string) error {
	newConfig := NewConfig()
	_, err := toml.DecodeFile(configFile, newConfig)
	if err != nil {
		return err
	}
	Conf = *newConfig

	return nil
}
