package responses

type MealIngredientResp struct {
	ID     string `json:"id" db:"id"`
	Name   string `json:"name" db:"ingredient_name"`
	Weight int    `json:"weight" db:"ingredient_weight"`
}
