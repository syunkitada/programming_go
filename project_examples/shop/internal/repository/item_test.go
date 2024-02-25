package repository_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/syunkitada/programming_go/project_examples/shop/internal/model"
	"github.com/syunkitada/programming_go/project_examples/shop/internal/repository"
)

func TestFindItems(t *testing.T) {
	t.Parallel()
	conf := repository.GetDefaultConfig()

	t.Run("find", func(t *testing.T) {
		t.Parallel()

		repo := repository.New(&conf)
		mock := repo.MustOpenMock()

		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "hoge")
		mock.ExpectQuery("^SELECT id,name FROM `items` WHERE deleted = 0$").WillReturnRows(rows)

		items, err := repo.FindItems(&repository.FindItemsInput{})
		assert.NoError(t, err)
		assert.Equal(t,
			[]model.Item{{ID: 1, Name: "hoge"}},
			items)
	})

	t.Run("find by id", func(t *testing.T) {
		t.Parallel()

		repo := repository.New(&conf)
		mock := repo.MustOpenMock()

		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "hoge")
		mock.ExpectQuery("^SELECT id,name FROM `items` WHERE deleted = 0 AND id = \\?$").
			WithArgs(1).
			WillReturnRows(rows)

		items, err := repo.FindItems(&repository.FindItemsInput{ID: 1})
		assert.NoError(t, err)
		assert.Equal(t,
			[]model.Item{{ID: 1, Name: "hoge"}},
			items)
	})
}

func TestItemSenario(t *testing.T) {
	t.Parallel()

	conf := repository.GetDefaultConfig()
	conf.Config.DBName = "TestItem"
	repo := repository.New(&conf)
	repo.MustRecreateDatabase()
	repo.MustOpen()
	repo.MustMigrate()

	var expectedItems []model.Item
	item1, err := repo.AddItem(&model.Item{Name: "hoge"})
	assert.NoError(t, err)

	expectedItems = append(expectedItems, *item1)
	items, err := repo.FindItems(&repository.FindItemsInput{})
	assert.NoError(t, err)
	assert.Equal(t, expectedItems, items)

	err = repo.DeleteItem(item1.ID)
	assert.NoError(t, err)

	items, err = repo.FindItems(&repository.FindItemsInput{})
	assert.NoError(t, err)
	assert.Equal(t, 0, len(items))
}
