// Code generated by Lingo for table sakila.film_actor - DO NOT EDIT

// +build !nolingo

package tfilmactor

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/path"
	"github.com/weworksandbox/lingo/sql"
)

func As(alias string) TFilmActor {
	return newTFilmActor(alias)
}

func New() TFilmActor {
	return newTFilmActor("")
}

func newTFilmActor(alias string) TFilmActor {
	t := TFilmActor{
		_alias: alias,
	}
	t.actorId = path.NewInt16(t, "actor_id")
	t.filmId = path.NewInt16(t, "film_id")
	t.lastUpdate = path.NewTime(t, "last_update")
	return t
}

type TFilmActor struct {
	_alias string

	actorId    path.Int16
	filmId     path.Int16
	lastUpdate path.Time
}

// lingo.Table Functions

func (t TFilmActor) GetColumns() []lingo.Column {
	return []lingo.Column{
		t.actorId,
		t.filmId,
		t.lastUpdate,
	}
}

func (t TFilmActor) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, t)
}

func (t TFilmActor) GetAlias() string {
	return t._alias
}

func (t TFilmActor) GetName() string {
	return "film_actor"
}

func (t TFilmActor) GetParent() string {
	return "sakila"
}

// Column Functions

func (t TFilmActor) ActorId() path.Int16 {
	return t.actorId
}

func (t TFilmActor) FilmId() path.Int16 {
	return t.filmId
}

func (t TFilmActor) LastUpdate() path.Time {
	return t.lastUpdate
}