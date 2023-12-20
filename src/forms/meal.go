package forms

type MealForm struct {
	Name        string `json:"name"`
	TotalWeight int    `json:"totalWeight"`
}

func (f MealForm) Validate() {
}

type MealIngredientForm struct {
	MealID       string `json:"mealId"`
	IngredientID string `json:"ingredientId"`
	TotalWeight  int    `json:"totalWeight"`
}

func (f MealIngredientForm) Validate() {
}
