package config

import (
	"github.com/urfave/cli"
)

type Config struct {
	App   *cli.App
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

func newConfig(ctx *cli.Context) *Config {
	defaultConfig := &Config{
		App: ctx.App,
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
