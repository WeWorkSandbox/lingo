// Code generated by Lingo for table sakila.film_actor - DO NOT EDIT

// +build !nolingo

package qfilmactor

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QFilmActor {
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
