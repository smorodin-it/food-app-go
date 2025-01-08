package responses

type MealIngredientResp struct {
	ID     string `json:"id" db:"id" validate:"required"`
	Name   string `json:"name" db:"ingredient_name" validate:"required"`
	Weight int    `json:"weight" db:"ingredient_weight" validate:"required"`
}
