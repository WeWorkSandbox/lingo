// Code generated by Lingo for table sakila.country - DO NOT EDIT

package qcountry

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func As(alias string) QCountry {
	return newQCountry(alias)
}

func New() QCountry {
	return newQCountry("")
}

func newQCountry(alias string) QCountry {
	q := QCountry{_alias: alias}
	q.countryId = path.NewInt16(q, "country_id")
	q.country = path.NewString(q, "country")
	q.lastUpdate = path.NewTime(q, "last_update")
	return q
}

type QCountry struct {
	_alias     string
	countryId  path.Int16
	country    path.String
	lastUpdate path.Time
}

// core.Table Functions

func (q QCountry) GetColumns() []core.Column {
	return []core.Column{
		q.countryId,
		q.country,
		q.lastUpdate,
	}
}

func (q QCountry) ToSQL(d core.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QCountry) GetAlias() string {
	return q._alias
}

func (q QCountry) GetName() string {
	return "country"
}

func (q QCountry) GetParent() string {
	return "sakila"
}

// Column Functions

func (q QCountry) CountryId() path.Int16 {
	return q.countryId
}

func (q QCountry) Country() path.String {
	return q.country
}

func (q QCountry) LastUpdate() path.Time {
	return q.lastUpdate
}
