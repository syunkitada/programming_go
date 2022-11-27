package models

type GetProject struct {
	Name string `validate:"required"`
}

type GetProjectData struct {
	Name string
}
