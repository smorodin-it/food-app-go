package repositories

import (
	"food-backend/src/database"
	"food-backend/src/domains"
	"food-backend/src/responses"
	"food-backend/src/utils"
	"time"
)

type MealRepository struct {
}

func (r MealRepository) List(page int, perPage int) (meals []domains.Meal, err error) {
	limit, offset, withPagination := utils.CalcPagination(page, perPage)

	if withPagination {
		sql := "SELECT * FROM meals ORDER BY created_at DESC LIMIT $1 OFFSET $2"
		err = database.DBCon.Select(&meals, sql, limit, offset)
	} else {
		sql := "SELECT * FROM meals ORDER BY created_at DESC"
		err = database.DBCon.Select(&meals, sql)
	}
	if err != nil {
		return nil, err
	}

	return meals, nil
}

func (r MealRepository) ListByUser(userId string, page int, perPage int) (meals []domains.Meal, err error) {
	limit, offset, withPagination := utils.CalcPagination(page, perPage)

	if withPagination {
		sql := "SELECT * FROM meals WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3"
		err = database.DBCon.Select(&meals, sql, userId, limit, offset)
	} else {
		sql := "SELECT * FROM meals WHERE user_id = $1 ORDER BY created_at DESC"
		err = database.DBCon.Select(&meals, sql, userId)
	}
	if err != nil {
		return nil, err
	}

	return meals, nil
}

func (r MealRepository) Create(meal domains.Meal) (id *string, err error) {
	sql := "INSERT INTO meals (meal_id, user_id, meal_name, total_weight) VALUES ($1, $2, $3, $4)"

	_, err = database.DBCon.Exec(sql, meal.MealID, meal.UserID, meal.MealName, meal.TotalWeight)
	if err != nil {
		return nil, err
	}

	return &meal.MealID, nil
}

func (r MealRepository) Retrieve(id string) (meal *domains.Meal, err error) {
	meal = &domains.Meal{}

	sql := "SELECT * FROM meals WHERE meal_id = $1"

	err = database.DBCon.Get(meal, sql, id)
	if err != nil {
		return nil, err
	}

	return meal, nil
}

func (r MealRepository) Update(meal domains.Meal, id string) (err error) {
	sql := "UPDATE meals SET meal_name = $1, total_weight = $2, created_at = $3 WHERE meal_id = $4"

	_, err = database.DBCon.Exec(sql, meal.MealName, meal.TotalWeight, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

func (r MealRepository) ListIngredients(mealId string) (ingredients *[]responses.MealIngredientResp, err error) {
	ingredients = new([]responses.MealIngredientResp)

	sql := "SELECT i.ingredient_id, i.ingredient_name, mi.ingredient_weight FROM meals_ingredients as mi INNER JOIN ingredients AS i ON mi.ingredient_id = i.ingredient_id WHERE mi.meal_id = $1"

	err = database.DBCon.Select(ingredients, sql, mealId)
	if err != nil {
		return nil, err
	}

	return ingredients, nil
}

func (r MealRepository) AddIngredient(mi domains.MealsIngredient) (id *string, err error) {
	sql := "INSERT INTO meals_ingredients (id, meal_id, ingredient_id, ingredient_weight) VALUES ($1, $2, $3, $4)"
	_, err = database.DBCon.Exec(sql, mi.ID, mi.MealId, mi.IngredientId, mi.IngredientWeight)
	if err != nil {
		return nil, err
	}

	return &mi.IngredientId, nil
}

func (r MealRepository) UpdateIngredient(mi domains.MealsIngredient) (err error) {
	sql := "UPDATE meals_ingredients SET ingredient_weight = $1, updated_at = $2 WHERE id = $3"
	_, err = database.DBCon.Exec(sql, mi.IngredientWeight, time.Now(), mi.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r MealRepository) DeleteIngredient(id string) (err error) {
	sql := "DELETE FROM meals_ingredients WHERE id = $1"
	_, err = database.DBCon.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}
