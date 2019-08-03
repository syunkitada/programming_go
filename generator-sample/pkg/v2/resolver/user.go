package resolver

import "github.com/syunkitada/go-samples/generator-sample/pkg/v2/models"

func (resolver *Resolver) GetUser(input *models.GetUser) (*models.GetUserData, int, error) {
	return &models.GetUserData{
		Name: "UserV2",
	}, 200, nil
}
