package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/syunkitada/programming_go/project_examples/shop/api/shop-api/oapi"
	"github.com/syunkitada/programming_go/project_examples/shop/internal/model"
	"github.com/syunkitada/programming_go/project_examples/shop/internal/repository"
)

func (self *Controller) FindItems(ctx echo.Context, params oapi.FindItemsParams) (items []oapi.Item, err error) {
	dbItems, err := self.repo.FindItems(&repository.FindItemsInput{})
	for _, dbItem := range dbItems {
		items = append(items, oapi.Item{
			Id:   dbItem.ID,
			Name: dbItem.Name,
		})
	}
	return items, nil
}

func (self *Controller) FindItemByID(ctx echo.Context, id uint64) (item oapi.Item, err error) {
	dbItems, err := self.repo.FindItems(&repository.FindItemsInput{ID: id})
	if err != nil {
		return item, err
	}
	if len(dbItems) > 1 {
		return item, echo.NewHTTPError(http.StatusConflict, fmt.Sprintf("multiple items found: id=%d", id))
	}
	if len(dbItems) == 0 {
		return item, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("item not found: id=%d", id))
	}
	dbItem := dbItems[0]
	item = oapi.Item{
		Id:   dbItem.ID,
		Name: dbItem.Name,
	}
	return item, nil
}

func (self *Controller) AddItem(ctx echo.Context, item *oapi.NewItem) error {
	if _, err := self.repo.AddItem(&model.Item{
		Name: item.Name,
	}); err != nil {
		return err
	}
	return nil
}

func (self *Controller) DeleteItem(ctx echo.Context, id uint64) error {
	if err := self.repo.DeleteItem(id); err != nil {
		return err
	}
	return nil
}
