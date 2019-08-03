package resolver

import "github.com/syunkitada/go-samples/generator-sample/pkg/v1/models"

func (resolver *Resolver) GetProject(input *models.GetProject) (*models.GetProjectData, int, error) {
	return &models.GetProjectData{
		Name: "Project",
	}, 200, nil
}
