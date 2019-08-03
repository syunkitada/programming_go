package models

type GetRole struct {
	Name string `validate:"required"`
}

type GetRoleData struct {
	Name string
}
