// Code generated by Lingo for table sakila.inventory - DO NOT EDIT

// +build !nolingo

package tinventory

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/path"
	"github.com/weworksandbox/lingo/sql"
)

func As(alias string) TInventory {
	return newTInventory(alias)
}

func New() TInventory {
	return newTInventory("")
}

func newTInventory(alias string) TInventory {
	t := TInventory{
		_alias: alias,
	}
	t.inventoryId = path.NewInt32(t, "inventory_id")
	t.filmId = path.NewInt16(t, "film_id")
	t.storeId = path.NewInt8(t, "store_id")
	t.lastUpdate = path.NewTime(t, "last_update")
	return t
}

type TInventory struct {
	_alias string

	inventoryId path.Int32
	filmId      path.Int16
	storeId     path.Int8
	lastUpdate  path.Time
}

// lingo.Table Functions

func (t TInventory) GetColumns() []lingo.Column {
	return []lingo.Column{
		t.inventoryId,
		t.filmId,
		t.storeId,
		t.lastUpdate,
	}
}

func (t TInventory) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, t)
}

func (t TInventory) GetAlias() string {
	return t._alias
}

func (t TInventory) GetName() string {
	return "inventory"
}

func (t TInventory) GetParent() string {
	return "sakila"
}

// Column Functions

func (t TInventory) InventoryId() path.Int32 {
	return t.inventoryId
}

func (t TInventory) FilmId() path.Int16 {
	return t.filmId
}

func (t TInventory) StoreId() path.Int8 {
	return t.storeId
}

func (t TInventory) LastUpdate() path.Time {
	return t.lastUpdate
}
