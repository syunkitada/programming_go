package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/syunkitada/programming_go/project_examples/shop/api/shop-api/oapi"
	"github.com/syunkitada/programming_go/project_examples/shop/internal/repository"
	"github.com/syunkitada/programming_go/project_examples/shop/internal/shop-api/config"
)

type IController interface {
	FindItems(ctx echo.Context, params oapi.FindItemsParams) (items []oapi.Item, err error)
	FindItemByID(ctx echo.Context, id uint64) (item oapi.Item, err error)
	AddItem(ctx echo.Context, item *oapi.NewItem) error
	DeleteItem(ctx echo.Context, id uint64) error
}

type Controller struct {
	conf *config.Config
	repo repository.IRepository
}

func New(conf *config.Config) IController {
	repo := repository.New(&conf.Repository)
	repo.MustOpen()

	return &Controller{
		conf: conf,
		repo: repo,
	}
}
