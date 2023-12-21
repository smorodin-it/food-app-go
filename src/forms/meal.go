package forms

type MealForm struct {
	Name        string `json:"name"`
	TotalWeight int    `json:"totalWeight"`
}

func (f MealForm) Validate() {
}

type MealIngredientAddForm struct {
	MealID       string `json:"mealId"`
	IngredientID string `json:"ingredientId"`
	TotalWeight  int    `json:"totalWeight"`
}

func (f MealIngredientAddForm) Validate() {
}

type MealIngredientUpdateForm struct {
	TotalWeight int `json:"totalWeight"`
}

func (f MealIngredientUpdateForm) Validate() {
}
