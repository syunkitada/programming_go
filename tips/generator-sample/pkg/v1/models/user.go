package models

type GetUser struct {
	Name string `validate:"required"`
}

type GetUserData struct {
	Name string
}

type PostUser struct {
	Name string `validate:"required"`
}

type PostUserData struct {
	Name string
}
