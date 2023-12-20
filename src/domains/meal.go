package domains

import "time"

type Meal struct {
	MealID      string    `json:"id" db:"meal_id"`
	UserID      string    `json:"-" db:"user_id"`
	MealName    string    `json:"name" db:"meal_name"`
	TotalWeight int       `json:"totalWeight" db:"total_weight"`
	CreatedAt   time.Time `json:"-" db:"created_at"`
	UpdatedAt   time.Time `json:"-" db:"updated_at"`
}

type MealsIngredient struct {
	ID               string    `json:"id" db:"id"`
	MealId           string    `json:"mealId" db:"meal_id"`
	IngredientId     string    `json:"ingredientId" db:"ingredient_id"`
	IngredientWeight int       `json:"ingredientWeight" db:"ingredient_weight"`
	CreatedAt        time.Time `json:"-" db:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"-" db:"updated_at" db:"updated_at"`
}
