// Code generated by Lingo for table sakila.customer - DO NOT EDIT

// +build !nolingo

package tcustomer

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/path"
	"github.com/weworksandbox/lingo/sql"
)

func As(alias string) TCustomer {
	return newTCustomer(alias)
}

func New() TCustomer {
	return newTCustomer("")
}

func newTCustomer(alias string) TCustomer {
	t := TCustomer{
		_alias: alias,
	}
	t.customerId = path.NewInt16(t, "customer_id")
	t.storeId = path.NewInt8(t, "store_id")
	t.firstName = path.NewString(t, "first_name")
	t.lastName = path.NewString(t, "last_name")
	t.email = path.NewString(t, "email")
	t.addressId = path.NewInt16(t, "address_id")
	t.active = path.NewInt8(t, "active")
	t.createDate = path.NewTime(t, "create_date")
	t.lastUpdate = path.NewTime(t, "last_update")
	return t
}

type TCustomer struct {
	_alias string

	customerId path.Int16
	storeId    path.Int8
	firstName  path.String
	lastName   path.String
	email      path.String
	addressId  path.Int16
	active     path.Int8
	createDate path.Time
	lastUpdate path.Time
}

// lingo.Table Functions

func (t TCustomer) GetColumns() []lingo.Column {
	return []lingo.Column{
		t.customerId,
		t.storeId,
		t.firstName,
		t.lastName,
		t.email,
		t.addressId,
		t.active,
		t.createDate,
		t.lastUpdate,
	}
}

func (t TCustomer) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, t)
}

func (t TCustomer) GetAlias() string {
	return t._alias
}

func (t TCustomer) GetName() string {
	return "customer"
}

func (t TCustomer) GetParent() string {
	return "sakila"
}

// Column Functions

func (t TCustomer) CustomerId() path.Int16 {
	return t.customerId
}

func (t TCustomer) StoreId() path.Int8 {
	return t.storeId
}

func (t TCustomer) FirstName() path.String {
	return t.firstName
}

func (t TCustomer) LastName() path.String {
	return t.lastName
}

func (t TCustomer) Email() path.String {
	return t.email
}

func (t TCustomer) AddressId() path.Int16 {
	return t.addressId
}

func (t TCustomer) Active() path.Int8 {
	return t.active
}

func (t TCustomer) CreateDate() path.Time {
	return t.createDate
}

func (t TCustomer) LastUpdate() path.Time {
	return t.lastUpdate
}
