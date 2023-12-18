package forms

type MealForm struct {
	Name        string `json:"name"`
	TotalWeight int    `json:"totalWeight"`
}

func (f MealForm) Validate() {
}
