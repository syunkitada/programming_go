package models

type GetUser struct {
	Name string `validate:"required"`
}

type GetUserData struct {
	Name string
}
