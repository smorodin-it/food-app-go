package repositories

import (
	"fmt"
	"food-backend/src/database"
	"food-backend/src/domains"
	"food-backend/src/utils"
)

type MealRepository struct {
}

func (r MealRepository) List(page int, perPage int) (meals []domains.Meal, err error) {
	limit, offset, withPagination := utils.CalcPagination(page, perPage)

	if withPagination {
		sql := "SELECT * FROM meals ORDER BY updated_at DESC LIMIT $1 OFFSET $2"
		err = database.DBCon.Select(&meals, sql, limit, offset)
	} else {
		sql := "SELECT * FROM meals ORDER BY updated_at DESC"
		err = database.DBCon.Select(&meals, sql)
	}
	if err != nil {
		return nil, err
	}

	return meals, nil
}

func (r MealRepository) Create(meal domains.Meal) (id *string, err error) {
	sql := "INSERT INTO meals (id, user_id, name, total_weight) VALUES ($1, $2, $3, $4)"

	result, err := database.DBCon.Exec(sql, meal.ID, meal.UserId, meal.Name, meal.TotalWeight)
	if err != nil {
		return nil, err
	}

	fmt.Println(result)

	return &meal.ID, nil

}

func (r MealRepository) Retrieve(id string) (meal *domains.Meal, err error) {
	meal = &domains.Meal{}

	sql := "SELECT * FROM meals WHERE id = $1"

	err = database.DBCon.Get(meal, sql, id)
	if err != nil {
		return nil, err
	}

	return meal, nil
}
