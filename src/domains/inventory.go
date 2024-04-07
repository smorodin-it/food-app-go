package domains

type Inventory struct {
	InventoryID string `json:"inventoryId" db:"inventory_id"`
	UserID      string `json:"userId" db:"user_id"`
	Name        string `json:"name" db:"name"`
	Weight      int    `json:"weight" db:"weight"`
}
