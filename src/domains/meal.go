package domains

import "time"

type Meal struct {
	MealID      string    `db:"meal_id" json:"id" validate:"required"`
	UserID      string    `db:"user_id" json:"-" validate:"required"`
	MealName    string    `db:"meal_name" json:"name" validate:"required"`
	TotalWeight int       `db:"total_weight" json:"totalWeight" validate:"required"`
	CreatedAt   time.Time `db:"created_at" json:"-" validate:"required"`
	UpdatedAt   time.Time `db:"updated_at" json:"-" validate:"required"`
}

type MealsIngredientAdd struct {
	ID               string    `db:"id" json:"id" validate:"required"`
	MealId           string    `db:"meal_id" json:"mealId" validate:"required"`
	IngredientId     string    `db:"ingredient_id" json:"ingredientId" validate:"required"`
	IngredientWeight int       `db:"ingredient_weight" json:"ingredientWeight" validate:"required"`
	CreatedAt        time.Time `db:"created_at" json:"-" validate:"required"`
	UpdatedAt        time.Time `db:"updated_at" json:"-" validate:"required"`
}

type MealsIngredientUpdate struct {
	ID               string `db:"id" json:"id" validate:"required"`
	IngredientWeight int    `db:"ingredient_weight" json:"ingredientWeight" validate:"required"`
}
