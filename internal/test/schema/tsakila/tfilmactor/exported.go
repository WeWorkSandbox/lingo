// Code generated by Lingo for table sakila.film_actor - DO NOT EDIT

// +build !nolingo

package tfilmactor

import (
	"github.com/weworksandbox/lingo/expr/path"
)

var instance = New()

func T() TFilmActor {
	return instance
}

func ActorId() path.Int16 {
	return instance.actorId
}

func FilmId() path.Int16 {
	return instance.filmId
}

func LastUpdate() path.Time {
	return instance.lastUpdate
}