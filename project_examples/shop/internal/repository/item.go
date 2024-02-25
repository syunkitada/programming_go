package repository

import (
	"github.com/syunkitada/programming_go/project_examples/shop/internal/model"
)

type FindItemsInput struct {
	ID   uint64
	Name string
}

func (self *Repository) FindItems(input *FindItemsInput) (items []model.Item, err error) {
	query := self.DB.Model(model.Item{}).
		Select("id,name").
		Where("deleted = 0")
	if input.ID != 0 {
		query.Where("id = ?", input.ID)
	}
	if input.Name != "" {
		query.Where("name = ?", input.Name)
	}

	if err = query.Scan(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (self *Repository) AddItem(item *model.Item) (*model.Item, error) {
	if err := self.DB.Model(model.Item{}).Save(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (self *Repository) DeleteItem(id uint64) (err error) {
	if err = self.DB.Where("id = ?", id).Delete(model.Item{}).Error; err != nil {
		return err
	}
	return nil
}
