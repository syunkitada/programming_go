package repository

import (
	"os"

	"github.com/syunkitada/programming_go/project_examples/shop/internal/model"
)

func (self *Repository) MustMigrate() {
	if err := self.Migrate(); err != nil {
		print("failed to self.Migrate", err.Error())
		os.Exit(1)
	}
}

func (self *Repository) Migrate() (err error) {
	if err = self.DB.AutoMigrate(
		&model.User{},
		&model.Item{},
	); err != nil {
		return err
	}
	return nil
}
