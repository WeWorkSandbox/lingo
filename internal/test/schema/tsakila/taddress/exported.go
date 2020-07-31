// Code generated by Lingo for table sakila.address - DO NOT EDIT

// +build !nolingo

package taddress

import (
	"github.com/weworksandbox/lingo/expr/path"
)

var instance = New()

func T() TAddress {
	return instance
}

func AddressId() path.Int16 {
	return instance.addressId
}

func Address() path.String {
	return instance.address
}

func Address2() path.String {
	return instance.address2
}

func District() path.String {
	return instance.district
}

func CityId() path.Int16 {
	return instance.cityId
}

func PostalCode() path.String {
	return instance.postalCode
}

func Phone() path.String {
	return instance.phone
}

func Location() path.Unsupported {
	return instance.location
}

func LastUpdate() path.Time {
	return instance.lastUpdate
}