package domains

import "time"

type Ingredient struct {
	IngredientID   string    `db:"ingredient_id" json:"id" validate:"required"`
	UserId         string    `db:"user_id" json:"-" validate:"required"`
	IngredientName string    `db:"ingredient_name" json:"name" validate:"required"`
	Manufacturer   string    `db:"manufacturer" json:"manufacturer" validate:"required"`
	Barcode        *string   `db:"barcode" json:"barcode"`
	Proteins       int       `db:"proteins" json:"proteins"`
	Carbs          int       `db:"carbs" json:"carbs"`
	Fats           int       `db:"fats" json:"fats"`
	Calories       int       `db:"calories" json:"calories"`
	CreatedAt      time.Time `db:"created_at" json:"-" validate:"required"`
	UpdatedAt      time.Time `db:"updated_at" json:"-" validate:"required"`
}
