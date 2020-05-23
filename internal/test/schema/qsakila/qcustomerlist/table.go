// Code generated by Lingo for table sakila.customer_list - DO NOT EDIT

package qcustomerlist

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QCustomerList {
	return newQCustomerList(alias)
}

func New() QCustomerList {
	return newQCustomerList("")
}

func newQCustomerList(alias string) QCustomerList {
	q := QCustomerList{_alias: alias}
	q.id = path.NewInt16(q, "ID")
	q.name = path.NewString(q, "name")
	q.address = path.NewString(q, "address")
	q.zipCode = path.NewString(q, "zip code")
	q.phone = path.NewString(q, "phone")
	q.city = path.NewString(q, "city")
	q.country = path.NewString(q, "country")
	q.notes = path.NewString(q, "notes")
	q.sid = path.NewInt8(q, "SID")
	return q
}

type QCustomerList struct {
	_alias  string
	id      path.Int16
	name    path.String
	address path.String
	zipCode path.String
	phone   path.String
	city    path.String
	country path.String
	notes   path.String
	sid     path.Int8
}

// core.Table Functions

func (q QCustomerList) GetColumns() []core.Column {
	return []core.Column{
		q.id,
		q.name,
		q.address,
		q.zipCode,
		q.phone,
		q.city,
		q.country,
		q.notes,
		q.sid,
	}
}

func (q QCustomerList) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QCustomerList) GetAlias() string {
	return q._alias
}

func (q QCustomerList) GetName() string {
	return "customer_list"
}

func (q QCustomerList) GetParent() string {
	return "sakila"
}

// Column Functions

func (q QCustomerList) Id() path.Int16 {
	return q.id
}

func (q QCustomerList) Name() path.String {
	return q.name
}

func (q QCustomerList) Address() path.String {
	return q.address
}

func (q QCustomerList) ZipCode() path.String {
	return q.zipCode
}

func (q QCustomerList) Phone() path.String {
	return q.phone
}

func (q QCustomerList) City() path.String {
	return q.city
}

func (q QCustomerList) Country() path.String {
	return q.country
}

func (q QCustomerList) Notes() path.String {
	return q.notes
}

func (q QCustomerList) Sid() path.Int8 {
	return q.sid
}
