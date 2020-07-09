// Code generated by Lingo for table sakila.film_list - DO NOT EDIT

// +build !nolingo

package qfilmlist

import "github.com/weworksandbox/lingo/pkg/core/expr/path"

var instance = New()

func Q() QFilmList {
	return instance
}

func Fid() path.Int16 {
	return instance.fid
}

func Title() path.String {
	return instance.title
}

func Description() path.String {
	return instance.description
}

func Category() path.String {
	return instance.category
}

func Price() path.Binary {
	return instance.price
}

func Length() path.Int16 {
	return instance.length
}

func Rating() path.String {
	return instance.rating
}

func Actors() path.String {
	return instance.actors
}
