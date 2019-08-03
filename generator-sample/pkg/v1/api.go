package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/syunkitada/go-samples/generator-sample/pkg/v1/models"
)

type Resolver interface {
	GetUser(input *models.GetUser) (*models.GetUserData, int, error)
	PostUser(input *models.PostUser) (*models.PostUserData, int, error)
	GetRole(input *models.GetRole) (*models.GetRoleData, int, error)
	GetProject(input *models.GetProject) (*models.GetProjectData, int, error)
}

func New(engine *gin.Engine, resolver Resolver) {
	handler := NewHandler(resolver)

	group := engine.Group("v1")
	{
		group.GET("/user", handler.GetUser)
		group.POST("/user", handler.PostUser)
		group.GET("/role", handler.GetRole)
		group.GET("/project", handler.GetProject)
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

func (router *Handler) PostUser(c *gin.Context) {
	data, statusCode, err := router.resolver.PostUser(&models.PostUser{})
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

func (router *Handler) GetRole(c *gin.Context) {
	data, statusCode, err := router.resolver.GetRole(&models.GetRole{})
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

func (router *Handler) GetProject(c *gin.Context) {
	data, statusCode, err := router.resolver.GetProject(&models.GetProject{})
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
