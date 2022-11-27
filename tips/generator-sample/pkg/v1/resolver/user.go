package resolver

import "github.com/syunkitada/go-samples/generator-sample/pkg/v1/models"

func (resolver *Resolver) GetUser(input *models.GetUser) (*models.GetUserData, int, error) {
	return &models.GetUserData{
		Name: "User",
	}, 200, nil
}

func (resolver *Resolver) PostUser(input *models.PostUser) (*models.PostUserData, int, error) {
	return &models.PostUserData{
		Name: "User",
	}, 200, nil
}
