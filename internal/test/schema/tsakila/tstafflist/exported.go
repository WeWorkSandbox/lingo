// Code generated by Lingo for table sakila.staff_list - DO NOT EDIT

// +build !nolingo

package tstafflist

import (
	"github.com/weworksandbox/lingo/expr/path"
)

var instance = New()

func T() TStaffList {
	return instance
}

func Id() path.Int8 {
	return instance.id
}

func Name() path.String {
	return instance.name
}

func Address() path.String {
	return instance.address
}

func ZipCode() path.String {
	return instance.zipCode
}

func Phone() path.String {
	return instance.phone
}

func City() path.String {
	return instance.city
}

func Country() path.String {
	return instance.country
}

func SID() path.Int8 {
	return instance.sID
}
