package db_models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string
	Password string
	Roles    []Role `gorm:"many2many:user_roles;"`
}
