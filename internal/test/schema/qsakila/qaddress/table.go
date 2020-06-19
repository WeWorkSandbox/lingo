// Code generated by Lingo for table sakila.address - DO NOT EDIT

// +build !nolingo

package qaddress

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func As(alias string) QAddress {
	return newQAddress(alias)
}

func New() QAddress {
	return newQAddress("")
}

func newQAddress(alias string) QAddress {
	q := QAddress{_alias: alias}
	q.addressId = path.NewInt16(q, "address_id")
	q.address = path.NewString(q, "address")
	q.address2 = path.NewString(q, "address2")
	q.district = path.NewString(q, "district")
	q.cityId = path.NewInt16(q, "city_id")
	q.postalCode = path.NewString(q, "postal_code")
	q.phone = path.NewString(q, "phone")
	q.location = path.NewUnsupported(q, "location")
	q.lastUpdate = path.NewTime(q, "last_update")
	return q
}

type QAddress struct {
	_alias     string
	addressId  path.Int16
	address    path.String
	address2   path.String
	district   path.String
	cityId     path.Int16
	postalCode path.String
	phone      path.String
	location   path.Unsupported
	lastUpdate path.Time
}

// core.Table Functions

func (q QAddress) GetColumns() []core.Column {
	return []core.Column{
		q.addressId,
		q.address,
		q.address2,
		q.district,
		q.cityId,
		q.postalCode,
		q.phone,
		q.location,
		q.lastUpdate,
	}
}

func (q QAddress) ToSQL(d core.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QAddress) GetAlias() string {
	return q._alias
}

func (q QAddress) GetName() string {
	return "address"
}

func (q QAddress) GetParent() string {
	return "sakila"
}

// Column Functions

func (q QAddress) AddressId() path.Int16 {
	return q.addressId
}

func (q QAddress) Address() path.String {
	return q.address
}

func (q QAddress) Address2() path.String {
	return q.address2
}

func (q QAddress) District() path.String {
	return q.district
}

func (q QAddress) CityId() path.Int16 {
	return q.cityId
}

func (q QAddress) PostalCode() path.String {
	return q.postalCode
}

func (q QAddress) Phone() path.String {
	return q.phone
}

func (q QAddress) Location() path.Unsupported {
	return q.location
}

func (q QAddress) LastUpdate() path.Time {
	return q.lastUpdate
}
