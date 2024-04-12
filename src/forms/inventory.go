package forms

type InventoryForm struct {
	UserID string `json:"userId" db:"user_id"`
	Name   string `json:"name" db:"name"`
	Weight int    `json:"weight" db:"weight"`
}
