// Code generated by Lingo for table sakila.inventory - DO NOT EDIT

// +build !nolingo

package qinventory

import "github.com/weworksandbox/lingo/expr/path"

var instance = New()

func Q() QInventory {
	return instance
}

func InventoryId() path.Int32 {
	return instance.inventoryId
}

func FilmId() path.Int16 {
	return instance.filmId
}

func StoreId() path.Int8 {
	return instance.storeId
}

func LastUpdate() path.Time {
	return instance.lastUpdate
}
