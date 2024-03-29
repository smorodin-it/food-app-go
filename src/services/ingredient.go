package services

import (
	"food-backend/src/domains"
	"food-backend/src/forms"
	"food-backend/src/repositories"
	"github.com/google/uuid"
)

type IngredientService interface {
	List(page int, perPage int) (ingredients []domains.Ingredient, err error)
	Retrieve(id string) (ingredient *domains.Ingredient, err error)
	Create(form *forms.IngredientForm, userId string) (id *string, err error)
	Update(form *forms.IngredientForm, id string) (err error)
}

type ingredientService struct {
	r repositories.IngredientRepository
}

func (s ingredientService) List(page int, perPage int) (ingredients []domains.Ingredient, err error) {
	ingredients, err = s.r.List(page, perPage)
	if err != nil {
		return nil, err
	}

	return ingredients, nil
}

func (s ingredientService) Retrieve(id string) (ingredient *domains.Ingredient, err error) {
	ingredient, err = s.r.Retrieve(id)
	if err != nil {
		return nil, err
	}

	return ingredient, nil
}

func (s ingredientService) Create(form *forms.IngredientForm, userId string) (id *string, err error) {
	model := domains.Ingredient{
		IngredientID:   uuid.New().String(),
		UserId:         userId,
		IngredientName: form.Name,
		Manufacturer:   form.Manufacturer,
		Barcode:        form.Barcode,
		Proteins:       form.Proteins,
		Carbs:          form.Carbs,
		Fats:           form.Fats,
		Calories:       form.Calories,
	}

	id, err = s.r.Create(&model, userId)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s ingredientService) Update(form *forms.IngredientForm, id string) (err error) {
	model := domains.Ingredient{
		IngredientID:   id,
		IngredientName: form.Name,
		Manufacturer:   form.Manufacturer,
		Barcode:        form.Barcode,
		Proteins:       form.Proteins,
		Carbs:          form.Carbs,
		Fats:           form.Fats,
		Calories:       form.Calories,
	}

	err = s.r.Update(&model)
	if err != nil {
		return err
	}

	return nil
}

func NewIngredientService(repository repositories.IngredientRepository) IngredientService {
	return &ingredientService{r: repository}
}
