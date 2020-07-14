// Code generated by Lingo for table sakila.sales_by_film_category - DO NOT EDIT

// +build !nolingo

package tsalesbyfilmcategory

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/path"
	"github.com/weworksandbox/lingo/sql"
)

func As(alias string) TSalesByFilmCategory {
	return newTSalesByFilmCategory(alias)
}

func New() TSalesByFilmCategory {
	return newTSalesByFilmCategory("")
}

func newTSalesByFilmCategory(alias string) TSalesByFilmCategory {
	t := TSalesByFilmCategory{
		_alias: alias,
	}
	t.category = path.NewString(t, "category")
	t.totalSales = path.NewBinary(t, "total_sales")
	return t
}

type TSalesByFilmCategory struct {
	_alias string

	category   path.String
	totalSales path.Binary
}

// lingo.Table Functions

func (t TSalesByFilmCategory) GetColumns() []lingo.Column {
	return []lingo.Column{
		t.category,
		t.totalSales,
	}
}

func (t TSalesByFilmCategory) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, t)
}

func (t TSalesByFilmCategory) GetAlias() string {
	return t._alias
}

func (t TSalesByFilmCategory) GetName() string {
	return "sales_by_film_category"
}

func (t TSalesByFilmCategory) GetParent() string {
	return "sakila"
}

// Column Functions

func (t TSalesByFilmCategory) Category() path.String {
	return t.category
}

func (t TSalesByFilmCategory) TotalSales() path.Binary {
	return t.totalSales
}
