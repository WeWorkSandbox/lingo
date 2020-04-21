// Code generated by Lingo for table sakila.customer_list - DO NOT EDIT

package qcustomerlist

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QCustomerList {
	return instance
}

func Id() path.Int16Path {
	return instance.id
}

func Name() path.StringPath {
	return instance.name
}

func Address() path.StringPath {
	return instance.address
}

func ZipCode() path.StringPath {
	return instance.zipCode
}

func Phone() path.StringPath {
	return instance.phone
}

func City() path.StringPath {
	return instance.city
}

func Country() path.StringPath {
	return instance.country
}

func Notes() path.StringPath {
	return instance.notes
}

func Sid() path.Int8Path {
	return instance.sid
}
