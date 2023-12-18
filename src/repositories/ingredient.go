package repositories

import (
	"food-backend/src/database"
	"food-backend/src/domains"
	"food-backend/src/utils"
	"time"
)

type IngredientRepository struct {
}

func (r IngredientRepository) List(page int, perPage int) (ingredients []domains.Ingredient, err error) {
	limit, offset, withPagination := utils.CalcPagination(page, perPage)

	if withPagination {
		sql := "SELECT * FROM ingredients ORDER BY updated_at DESC LIMIT $1 OFFSET $2"
		err = database.DBCon.Select(&ingredients, sql, limit, offset)
	} else {
		sql := "SELECT * FROM ingredients ORDER BY updated_at DESC"
		err = database.DBCon.Select(&ingredients, sql)
	}

	if err != nil {
		return nil, err
	}

	return ingredients, nil
}

func (r IngredientRepository) Create(ingredient *domains.Ingredient, userId string) (*string, error) {
	sql := "INSERT INTO ingredients (id, user_id, name, manufacturer, barcode, proteins, carbs, fats, calories) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	_, err := database.DBCon.Exec(sql, ingredient.ID, userId, ingredient.Name, ingredient.Manufacturer, ingredient.Barcode,
		ingredient.Proteins, ingredient.Carbs, ingredient.Fats, ingredient.Calories)
	if err != nil {
		return nil, err
	}

	return &ingredient.ID, nil
}

func (r IngredientRepository) Update(ingredient *domains.Ingredient) error {
	sql := "UPDATE ingredients SET  name = $1, manufacturer = $2, barcode = $3, proteins = $4, carbs = $5, fats = $6, calories = $7, updated_at = $8 WHERE id = $9"

	_, err := database.DBCon.Exec(sql, ingredient.Name, ingredient.Manufacturer, ingredient.Barcode, ingredient.Proteins, ingredient.Carbs, ingredient.Fats, ingredient.Calories, time.Now(), ingredient.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r IngredientRepository) Retrieve(id string) (*domains.Ingredient, error) {
	model := new(domains.Ingredient)

	sql := "SELECT * FROM ingredients WHERE id = $1"

	err := database.DBCon.Get(model, sql, id)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r IngredientRepository) GetByBarcode(barcode *string) (*domains.Ingredient, error) {
	if barcode != nil {
		model := new(domains.Ingredient)

		sql := "SELECT * FROM ingredients WHERE barcode = $1"
		err := database.DBCon.Get(model, sql, barcode)
		if err != nil {
			return nil, err
		}

		return model, nil
	}

	return nil, nil
}
