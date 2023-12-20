package services

import (
	"food-backend/src/domains"
	"food-backend/src/forms"
	"food-backend/src/repositories"
	"github.com/google/uuid"
	"time"
)

type MealService struct {
	r repositories.MealRepository
}

func (s MealService) List(page int, perPage int) (meals []domains.Meal, err error) {
	meals, err = s.r.List(page, perPage)
	if err != nil {
		return nil, err
	}

	return meals, nil
}

func (s MealService) Create(f forms.MealForm, userId string) (id *string, err error) {
	m := domains.Meal{
		MealID:      uuid.New().String(),
		UserID:      userId,
		MealName:    f.Name,
		TotalWeight: f.TotalWeight,
	}

	id, err = s.r.Create(m)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s MealService) Retrieve(id string) (meal *domains.Meal, err error) {
	meal, err = s.r.Retrieve(id)
	if err != nil {
		return nil, err
	}

	return meal, nil

}

func (s MealService) Update(f forms.MealForm, id string) (err error) {
	m := domains.Meal{
		MealName:    f.Name,
		TotalWeight: f.TotalWeight,
		UpdatedAt:   time.Now(),
	}

	err = s.r.Update(m, id)
	if err != nil {
		return err
	}

	return nil
}
