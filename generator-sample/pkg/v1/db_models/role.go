package db_models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name string
}
