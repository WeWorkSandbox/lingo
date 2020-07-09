// Code generated by Lingo for table sakila.film_actor - DO NOT EDIT

// +build !nolingo

package qfilmactor

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/path"
	"github.com/weworksandbox/lingo/sql"
)

func As(alias string) QFilmActor {
	return newQFilmActor(alias)
}

func New() QFilmActor {
	return newQFilmActor("")
}

func newQFilmActor(alias string) QFilmActor {
	q := QFilmActor{_alias: alias}
	q.actorId = path.NewInt16(q, "actor_id")
	q.filmId = path.NewInt16(q, "film_id")
	q.lastUpdate = path.NewTime(q, "last_update")
	return q
}

type QFilmActor struct {
	_alias     string
	actorId    path.Int16
	filmId     path.Int16
	lastUpdate path.Time
}

// lingo.Table Functions

func (q QFilmActor) GetColumns() []lingo.Column {
	return []lingo.Column{
		q.actorId,
		q.filmId,
		q.lastUpdate,
	}
}

func (q QFilmActor) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QFilmActor) GetAlias() string {
	return q._alias
}

func (q QFilmActor) GetName() string {
	return "film_actor"
}

func (q QFilmActor) GetParent() string {
	return "sakila"
}

// Column Functions

func (q QFilmActor) ActorId() path.Int16 {
	return q.actorId
}

func (q QFilmActor) FilmId() path.Int16 {
	return q.filmId
}

func (q QFilmActor) LastUpdate() path.Time {
	return q.lastUpdate
}
