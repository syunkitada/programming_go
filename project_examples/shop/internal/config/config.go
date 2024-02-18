package config

import (
	"github.com/syunkitada/programming_go/project_examples/shop/internal/repository"
)

type Config struct {
	Repository repository.Config
}

func GetDefaultConfig() Config {
	return Config{
		Repository: repository.GetDefaultConfig(),
	}
}
