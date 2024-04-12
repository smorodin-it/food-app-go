package repositories

import (
	"food-backend/src/domains"
	"food-backend/src/utils"
	"github.com/jmoiron/sqlx"
)

type InventoryRepo interface {
	List(page int, perPage int) (inventories []domains.Inventory, err error)
	Create(inventory *domains.Inventory) (id *string, err error)
	Update(inventory *domains.Inventory) (err error)
	Retrieve(id string) (inventory *domains.Inventory, err error)
}

type inventoryRepo struct {
	db *sqlx.DB
}

func (r *inventoryRepo) List(page int, perPage int) (inventories []domains.Inventory, err error) {
	limit, offset, withPagination := utils.CalcPagination(page, perPage)

	if withPagination {
		sql := "SELECT * FROM inventory ORDER BY name LIMIT $1? OFFSET $2"
		err = r.db.Select(&inventories, sql, limit, offset)
	} else {
		sql := "SELECT * FROM inventory ORDER BY name"
		err = r.db.Select(&inventories, sql)
	}

	if err != nil {
		return nil, err
	}

	return inventories, nil
}

func (r *inventoryRepo) Create(inventory *domains.Inventory) (id *string, err error) {
	sql := "INSERT INTO inventory(inventory_id, user_id, name, weight) VALUES ($1, $2, $3, $4)"

	_, err = r.db.Exec(sql, inventory.InventoryID, inventory.UserID, inventory.Name, inventory.Weight)
	if err != nil {
		return nil, err
	}

	return &inventory.InventoryID, nil
}

func (r *inventoryRepo) Update(inventory *domains.Inventory) (err error) {
	sql := "UPDATE inventory SET name = $1, weight = $2 WHERE inventory_id = $3"
	_, err = r.db.Exec(sql, inventory.Name, inventory.Weight, inventory.InventoryID)

	if err != nil {
		return err
	}

	return nil
}

func (r *inventoryRepo) Retrieve(id string) (inventory *domains.Inventory, err error) {
	model := new(domains.Inventory)

	sql := "SELECT * FROM inventory WHERE inventory_id = $1"
	err = r.db.Get(model, sql, id)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func NewInventoryRepo(db *sqlx.DB) InventoryRepo { return &inventoryRepo{db: db} }
