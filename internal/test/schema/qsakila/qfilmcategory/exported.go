// Code generated by Lingo for table sakila.film_category - DO NOT EDIT

// +build !nolingo

package qfilmcategory

import (
	"github.com/weworksandbox/lingo/expr/path"
)

var instance = New()

func Q() QFilmCategory {
	return instance
}

func FilmId() path.Int16 {
	return instance.filmId
}

func CategoryId() path.Int8 {
	return instance.categoryId
}

func LastUpdate() path.Time {
	return instance.lastUpdate
}
