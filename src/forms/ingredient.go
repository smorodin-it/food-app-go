package forms

type IngredientForm struct {
	Name         string  `json:"name" validate:"required"`
	Manufacturer string  `json:"manufacturer" validate:"required"`
	Barcode      *string `json:"barcode"`
	Proteins     int     `json:"proteins"`
	Carbs        int     `json:"carbs"`
	Fats         int     `json:"fats"`
	Calories     int     `json:"calories"`
}
