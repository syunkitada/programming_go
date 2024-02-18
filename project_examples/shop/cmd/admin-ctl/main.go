package main

import (
	"github.com/syunkitada/programming_go/project_examples/shop/internal/config"
	"github.com/syunkitada/programming_go/project_examples/shop/internal/repository"
)

func main() {
	conf := config.GetDefaultConfig()
	repo := repository.New(&conf.Repository)
	repo.MustCreateDatabase()
	repo.MustOpen()
	repo.MustMigrate()
}
