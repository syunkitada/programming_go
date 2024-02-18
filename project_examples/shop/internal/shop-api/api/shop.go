package api

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/programming_go/project_examples/shop/api/shop-api/oapi"
	"github.com/syunkitada/programming_go/project_examples/shop/internal/repository"
	"github.com/syunkitada/programming_go/project_examples/shop/internal/shop-api/config"
)

type Shop struct {
	Pets   map[int64]oapi.Pet
	NextId int64
	Lock   sync.Mutex
}

func NewShop(conf *config.Config) *Shop {
	repo := repository.New(&conf.Repository)
	repo.MustOpen()

	return &Shop{
		Pets:   make(map[int64]oapi.Pet),
		NextId: 1000,
	}
}

// sendShopError wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendShopError(ctx echo.Context, code int, message string) error {
	petErr := oapi.Error{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, petErr)
	return err
}

// FindPets implements all the handlers in the ServerInterface
func (p *Shop) FindPets(ctx echo.Context, params oapi.FindPetsParams) error {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	var result []oapi.Pet

	for _, pet := range p.Pets {
		if params.Tags != nil {
			// If we have tags,  filter pets by tag
			for _, t := range *params.Tags {
				if pet.Tag != nil && (*pet.Tag == t) {
					result = append(result, pet)
				}
			}
		} else {
			// Add all pets if we're not filtering
			result = append(result, pet)
		}

		if params.Limit != nil {
			l := int(*params.Limit)
			if len(result) >= l {
				// We're at the limit
				break
			}
		}
	}
	return ctx.JSON(http.StatusOK, result)
}

func (p *Shop) AddPet(ctx echo.Context) error {
	// We expect a NewPet object in the request body.
	var newPet oapi.NewPet
	err := ctx.Bind(&newPet)
	if err != nil {
		return sendShopError(ctx, http.StatusBadRequest, "Invalid format for NewPet")
	}
	// We now have a pet, let's add it to our "database".

	// We're always asynchronous, so lock unsafe operations below
	p.Lock.Lock()
	defer p.Lock.Unlock()

	// We handle pets, not NewPets, which have an additional ID field
	var pet oapi.Pet
	pet.Name = newPet.Name
	pet.Tag = newPet.Tag
	pet.Id = p.NextId
	p.NextId++

	// Insert into map
	p.Pets[pet.Id] = pet

	// Now, we have to return the NewPet
	err = ctx.JSON(http.StatusCreated, pet)
	if err != nil {
		// Something really bad happened, tell Echo that our handler failed
		return err
	}

	// Return no error. This refers to the handler. Even if we return an HTTP
	// error, but everything else is working properly, tell Echo that we serviced
	// the error. We should only return errors from Echo handlers if the actual
	// servicing of the error on the infrastructure level failed. Returning an
	// HTTP/400 or HTTP/500 from here means Echo/HTTP are still working, so
	// return nil.
	return nil
}

func (p *Shop) FindPetByID(ctx echo.Context, petId int64) error {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	pet, found := p.Pets[petId]
	if !found {
		return sendShopError(ctx, http.StatusNotFound,
			fmt.Sprintf("Could not find pet with ID %d", petId))
	}
	return ctx.JSON(http.StatusOK, pet)
}

func (p *Shop) DeletePet(ctx echo.Context, id int64) error {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	_, found := p.Pets[id]
	if !found {
		return sendShopError(ctx, http.StatusNotFound,
			fmt.Sprintf("Could not find pet with ID %d", id))
	}
	delete(p.Pets, id)
	return ctx.NoContent(http.StatusNoContent)
}
