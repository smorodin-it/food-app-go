package services

import (
	"food-backend/src/domains"
	"food-backend/src/forms"
	"food-backend/src/repositories"
	"github.com/google/uuid"
)

type InventoryService interface {
	List(page int, perPage int) (inventories []domains.Inventory, err error)
	Retrieve(id string) (inventory *domains.Inventory, err error)
	Create(form *forms.InventoryForm, userId string) (id *string, err error)
	Update(form *forms.InventoryForm, id string) (err error)
}

type inventoryService struct {
	r repositories.InventoryRepo
}

func (s inventoryService) List(page int, perPage int) (inventories []domains.Inventory, err error) {
	inventories, err = s.r.List(page, perPage)
	if err != nil {
		return nil, err
	}

	return inventories, nil
}

func (s inventoryService) Retrieve(id string) (inventory *domains.Inventory, err error) {
	inventory, err = s.r.Retrieve(id)
	if err != nil {
		return nil, err
	}

	return inventory, nil
}

func (s inventoryService) Create(form *forms.InventoryForm, userId string) (id *string, err error) {
	model := domains.Inventory{
		InventoryID: uuid.New().String(),
		UserID:      userId,
		Name:        form.Name,
		Weight:      form.Weight,
	}

	id, err = s.r.Create(&model)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s inventoryService) Update(form *forms.InventoryForm, id string) (err error) {
	model := domains.Inventory{
		InventoryID: id,
		UserID:      form.UserID,
		Name:        form.Name,
		Weight:      form.Weight,
	}

	err = s.r.Update(&model)
	if err != nil {
		return err
	}
	return nil
}

func NewInventoryService(repository repositories.InventoryRepo) InventoryService {
	return &inventoryService{r: repository}
}
