// Code generated by Lingo for table sakila.actor_info - DO NOT EDIT

// +build !nolingo

package qactorinfo

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func As(alias string) QActorInfo {
	return newQActorInfo(alias)
}

func New() QActorInfo {
	return newQActorInfo("")
}

func newQActorInfo(alias string) QActorInfo {
	q := QActorInfo{_alias: alias}
	q.actorId = path.NewInt16(q, "actor_id")
	q.firstName = path.NewString(q, "first_name")
	q.lastName = path.NewString(q, "last_name")
	q.filmInfo = path.NewString(q, "film_info")
	return q
}

type QActorInfo struct {
	_alias    string
	actorId   path.Int16
	firstName path.String
	lastName  path.String
	filmInfo  path.String
}

// core.Table Functions

func (q QActorInfo) GetColumns() []core.Column {
	return []core.Column{
		q.actorId,
		q.firstName,
		q.lastName,
		q.filmInfo,
	}
}

func (q QActorInfo) ToSQL(d core.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QActorInfo) GetAlias() string {
	return q._alias
}

func (q QActorInfo) GetName() string {
	return "actor_info"
}

func (q QActorInfo) GetParent() string {
	return "sakila"
}

// Column Functions

func (q QActorInfo) ActorId() path.Int16 {
	return q.actorId
}

func (q QActorInfo) FirstName() path.String {
	return q.firstName
}

func (q QActorInfo) LastName() path.String {
	return q.lastName
}

func (q QActorInfo) FilmInfo() path.String {
	return q.filmInfo
}
