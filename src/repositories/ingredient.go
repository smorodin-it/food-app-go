package repositories

import (
	"food-backend/src/domains"
	"food-backend/src/utils"
	"github.com/jmoiron/sqlx"
	"time"
)

type IngredientRepository interface {
	List(page int, perPage int) (ingredients []domains.Ingredient, err error)
	Create(ingredient *domains.Ingredient, userId string) (*string, error)
	Update(ingredient *domains.Ingredient) error
	Retrieve(id string) (*domains.Ingredient, error)
	GetByBarcode(barcode *string) (*domains.Ingredient, error)
}

type ingredientRepo struct {
	db *sqlx.DB
}

func (r ingredientRepo) List(page int, perPage int) (ingredients []domains.Ingredient, err error) {
	limit, offset, withPagination := utils.CalcPagination(page, perPage)

	if withPagination {
		sql := "SELECT * FROM ingredients ORDER BY updated_at DESC LIMIT $1 OFFSET $2"
		err = r.db.Select(&ingredients, sql, limit, offset)
	} else {
		sql := "SELECT * FROM ingredients ORDER BY updated_at DESC"
		err = r.db.Select(&ingredients, sql)
	}

	if err != nil {
		return nil, err
	}

	return ingredients, nil
}

func (r ingredientRepo) Create(ingredient *domains.Ingredient, userId string) (*string, error) {
	sql := "INSERT INTO ingredients (ingredient_id, user_id, ingredient_name, manufacturer, barcode, proteins, carbs, fats, calories) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	_, err := r.db.Exec(sql, ingredient.IngredientID, userId, ingredient.IngredientName, ingredient.Manufacturer, ingredient.Barcode,
		ingredient.Proteins, ingredient.Carbs, ingredient.Fats, ingredient.Calories)
	if err != nil {
		return nil, err
	}

	return &ingredient.IngredientID, nil
}

func (r ingredientRepo) Update(ingredient *domains.Ingredient) error {
	sql := "UPDATE ingredients SET  ingredient_name = $1, manufacturer = $2, barcode = $3, proteins = $4, carbs = $5, fats = $6, calories = $7, updated_at = $8 WHERE ingredient_id = $9"

	_, err := r.db.Exec(sql, ingredient.IngredientName, ingredient.Manufacturer, ingredient.Barcode, ingredient.Proteins, ingredient.Carbs, ingredient.Fats, ingredient.Calories, time.Now(), ingredient.IngredientID)
	if err != nil {
		return err
	}

	return nil
}

func (r ingredientRepo) Retrieve(id string) (*domains.Ingredient, error) {
	model := new(domains.Ingredient)

	sql := "SELECT * FROM ingredients WHERE ingredient_id = $1"

	err := r.db.Get(model, sql, id)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r ingredientRepo) GetByBarcode(barcode *string) (*domains.Ingredient, error) {
	if barcode != nil {
		model := new(domains.Ingredient)

		sql := "SELECT * FROM ingredients WHERE barcode = $1"
		err := r.db.Get(model, sql, barcode)
		if err != nil {
			return nil, err
		}

		return model, nil
	}

	return nil, nil
}

func NewIngredientRepo(db *sqlx.DB) IngredientRepository {
	return &ingredientRepo{db: db}
}
