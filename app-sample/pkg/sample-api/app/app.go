package app

import (
	"sync"
	"sync/atomic"

	"github.com/go-openapi/errors"

	"github.com/syunkitada/programming_go/app-sample/pkg/sample-api/config"
	"github.com/syunkitada/programming_go/app-sample/pkg/sample-api/gen/models"
)

type App struct {
}

var lastID int64
var itemsLock = &sync.Mutex{}
var items = map[int64]*models.Item{}

func New(conf *config.Config) *App {
	return &App{}
}

func newItemID() int64 {
	return atomic.AddInt64(&lastID, 1)
}

func (self *App) AddItem(item *models.Item) error {
	if item == nil {
		return errors.New(500, "item must be present")
	}

	itemsLock.Lock()
	defer itemsLock.Unlock()

	newID := newItemID()
	item.ID = newID
	items[newID] = item

	return nil
}

func (self *App) UpdateItem(id int64, item *models.Item) error {
	if item == nil {
		return errors.New(500, "item must be present")
	}

	itemsLock.Lock()
	defer itemsLock.Unlock()

	_, exists := items[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	item.ID = id
	items[id] = item
	return nil
}

func (self *App) DeleteItem(id int64) error {
	itemsLock.Lock()
	defer itemsLock.Unlock()

	_, exists := items[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	delete(items, id)
	return nil
}

func (self *App) AllItems(since int64, limit int32) (result []*models.Item) {
	result = make([]*models.Item, 0)
	for id, item := range items {
		if len(result) >= int(limit) {
			return
		}
		if since == 0 || id > since {
			result = append(result, item)
		}
	}
	return
}
