// Code generated by Lingo for table sakila.film_text - DO NOT EDIT

// +build !nolingo

package qfilmtext

import "github.com/weworksandbox/lingo/pkg/core/expr/path"

var instance = New()

func Q() QFilmText {
	return instance
}

func FilmId() path.Int16 {
	return instance.filmId
}

func Title() path.String {
	return instance.title
}

func Description() path.String {
	return instance.description
}
