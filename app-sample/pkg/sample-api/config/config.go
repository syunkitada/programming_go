package config

import (
	"github.com/syunkitada/programming_go/app-sample/pkg/lib/logger"
)

type Config struct {
	Logger logger.Config
}

var DefaultConfig = Config{
	Logger: logger.Config{
		OutputPaths: []string{"stdout"},
		Level:       "info",
		Encoding:    "json",
	},
}

func (self *Config) Complete() (err error) {
	return
}

func (self *Config) Validate() (err error) {
	return
}
