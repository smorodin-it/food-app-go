package domains

import "time"

type Ingredient struct {
	IngredientID   string    `json:"id" db:"ingredient_id"`
	UserId         string    `json:"-" db:"user_id"`
	IngredientName string    `json:"name" db:"ingredient_name"`
	Manufacturer   string    `json:"manufacturer" db:"manufacturer"`
	Barcode        *string   `json:"barcode" db:"barcode"`
	Proteins       int       `json:"proteins" db:"proteins"`
	Carbs          int       `json:"carbs" db:"carbs"`
	Fats           int       `json:"fats" db:"fats"`
	Calories       int       `json:"calories" db:"calories"`
	CreatedAt      time.Time `json:"-" db:"created_at"`
	UpdatedAt      time.Time `json:"-" db:"updated_at"`
}
