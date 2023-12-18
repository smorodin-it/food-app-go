package services

import (
	"food-backend/src/domains"
	"food-backend/src/forms"
	"food-backend/src/repositories"
	"github.com/google/uuid"
)

type IngredientService struct {
	r repositories.IngredientRepository
}

func (s IngredientService) List(page int, perPage int) (ingredients []domains.Ingredient, err error) {
	ingredients, err = s.r.List(page, perPage)
	if err != nil {
		return nil, err
	}

	return ingredients, nil
}

func (s IngredientService) Retrieve(id string) (ingredient *domains.Ingredient, err error) {
	ingredient, err = s.r.Retrieve(id)
	if err != nil {
		return nil, err
	}

	return ingredient, nil
}

func (s IngredientService) Create(form *forms.IngredientForm, userId string) (id *string, err error) {
	model := domains.Ingredient{
		ID:           uuid.New().String(),
		UserId:       userId,
		Name:         form.Name,
		Manufacturer: form.Manufacturer,
		Barcode:      form.Barcode,
		Proteins:     form.Proteins,
		Carbs:        form.Carbs,
		Fats:         form.Fats,
		Calories:     form.Calories,
	}

	id, err = s.r.Create(&model, userId)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s IngredientService) Update(form *forms.IngredientForm, id string) (err error) {
	model := domains.Ingredient{
		ID:           id,
		Name:         form.Name,
		Manufacturer: form.Manufacturer,
		Barcode:      form.Barcode,
		Proteins:     form.Proteins,
		Carbs:        form.Carbs,
		Fats:         form.Fats,
		Calories:     form.Calories,
	}

	err = s.r.Update(&model)
	if err != nil {
		return err
	}

	return nil
}
