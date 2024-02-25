package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/programming_go/project_examples/shop/api/shop-api/oapi"
	"github.com/syunkitada/programming_go/project_examples/shop/internal/shop-api/config"
	"github.com/syunkitada/programming_go/project_examples/shop/internal/shop-api/controller"
)

type Handler struct {
	ctrl controller.IController
}

func NewHandler(conf *config.Config) *Handler {
	ctrl := controller.New(conf)

	return &Handler{
		ctrl: ctrl,
	}
}

// sendHandlerError wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendHandlerError(ctx echo.Context, code int, message string) error {
	itemErr := oapi.Error{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, itemErr)
	return err
}

func (self *Handler) FindItems(ctx echo.Context, params oapi.FindItemsParams) error {
	items, err := self.ctrl.FindItems(ctx, params)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, items)
}

func (self *Handler) AddItem(ctx echo.Context) error {
	var newItem oapi.NewItem
	err := ctx.Bind(&newItem)
	if err != nil {
		return sendHandlerError(ctx, http.StatusBadRequest, "Invalid format for NewItem")
	}

	if err := self.ctrl.AddItem(ctx, &newItem); err != nil {
		return err
	}
	return nil
}

func (self *Handler) FindItemByID(ctx echo.Context, itemId uint64) error {
	item, err := self.ctrl.FindItemByID(ctx, itemId)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, item)
}

func (self *Handler) DeleteItem(ctx echo.Context, id uint64) error {
	err := self.ctrl.DeleteItem(ctx, id)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent)
}
