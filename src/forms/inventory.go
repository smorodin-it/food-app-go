package forms

type InventoryForm struct {
	Name   string `json:"name" validate:"required"`
	Weight int    `json:"weight" validate:"required"`
}
