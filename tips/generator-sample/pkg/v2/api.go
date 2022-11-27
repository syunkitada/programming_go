package v2

import (
	"github.com/gin-gonic/gin"

	"github.com/syunkitada/go-samples/generator-sample/pkg/v2/models"
)

type Resolver interface {
	GetUser(input *models.GetUser) (*models.GetUserData, int, error)
}

func New(engine *gin.Engine, resolver Resolver) {
	handler := NewHandler(resolver)

	group := engine.Group("v2")
	{
		group.GET("/user", handler.GetUser)
	}
}

type Handler struct {
	resolver Resolver
}

func NewHandler(resolver Resolver) *Handler {
	return &Handler{resolver: resolver}
}

func (router *Handler) GetUser(c *gin.Context) {
	data, statusCode, err := router.resolver.GetUser(&models.GetUser{})
	if err != nil {
		c.JSON(statusCode, gin.H{
			"Data":  data,
			"Error": err,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"Data": data,
	})
}
