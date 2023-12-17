package forms

type IngredientForm struct {
	Name         string  `json:"name"`
	Manufacturer string  `json:"manufacturer"`
	Barcode      *string `json:"barcode"`
	Proteins     string  `json:"proteins"`
	Carbs        string  `json:"carbs"`
	Fats         string  `json:"fats"`
	Calories     string  `json:"calories"`
}
