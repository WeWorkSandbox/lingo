// Code generated by Lingo for table sakila.city - DO NOT EDIT

// +build !nolingo

package qcity

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expr/path"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func As(alias string) QCity {
	return newQCity(alias)
}

func New() QCity {
	return newQCity("")
}

func newQCity(alias string) QCity {
	q := QCity{_alias: alias}
	q.cityId = path.NewInt16(q, "city_id")
	q.city = path.NewString(q, "city")
	q.countryId = path.NewInt16(q, "country_id")
	q.lastUpdate = path.NewTime(q, "last_update")
	return q
}

type QCity struct {
	_alias     string
	cityId     path.Int16
	city       path.String
	countryId  path.Int16
	lastUpdate path.Time
}

// core.Table Functions

func (q QCity) GetColumns() []core.Column {
	return []core.Column{
		q.cityId,
		q.city,
		q.countryId,
		q.lastUpdate,
	}
}

func (q QCity) ToSQL(d core.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QCity) GetAlias() string {
	return q._alias
}

func (q QCity) GetName() string {
	return "city"
}

func (q QCity) GetParent() string {
	return "sakila"
}

// Column Functions

func (q QCity) CityId() path.Int16 {
	return q.cityId
}

func (q QCity) City() path.String {
	return q.city
}

func (q QCity) CountryId() path.Int16 {
	return q.countryId
}

func (q QCity) LastUpdate() path.Time {
	return q.lastUpdate
}
