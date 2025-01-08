package forms

type MealForm struct {
	Name        string `json:"name" validate:"required"`
	TotalWeight int    `json:"totalWeight" validate:"required"`
}

func (f MealForm) Validate() {
}

type MealIngredientAddForm struct {
	MealID       string `json:"mealId" validate:"required"`
	IngredientID string `json:"ingredientId" validate:"required"`
	TotalWeight  int    `json:"totalWeight" validate:"required"`
}

func (f MealIngredientAddForm) Validate() {
}

type MealIngredientUpdateForm struct {
	TotalWeight int `json:"totalWeight" validate:"required"`
}

func (f MealIngredientUpdateForm) Validate() {
}
