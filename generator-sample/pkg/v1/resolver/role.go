package resolver

import "github.com/syunkitada/go-samples/generator-sample/pkg/v1/models"

func (resolver *Resolver) GetRole(input *models.GetRole) (*models.GetRoleData, int, error) {
	return &models.GetRoleData{
		Name: "Role",
	}, 200, nil
}
