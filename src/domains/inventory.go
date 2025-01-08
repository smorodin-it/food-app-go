package domains

import "time"

type Inventory struct {
	InventoryID string    `db:"inventory_id" json:"id" validate:"required"`
	UserID      string    `db:"user_id" json:"-" validate:"required"`
	Name        string    `db:"name" json:"name" validate:"required"`
	Weight      int       `db:"weight" json:"weight" validate:"required"`
	CreatedAt   time.Time `db:"created_at" json:"-" validate:"required"`
	UpdatedAt   time.Time `db:"updated_at" json:"-" validate:"required"`
}
