// Code generated by Lingo for table sakila.nicer_but_slower_film_list - DO NOT EDIT

package qnicerbutslowerfilmlist

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QNicerButSlowerFilmList {
	return newQNicerButSlowerFilmList(alias)
}

func New() QNicerButSlowerFilmList {
	return newQNicerButSlowerFilmList("")
}

func newQNicerButSlowerFilmList(alias string) QNicerButSlowerFilmList {
	q := QNicerButSlowerFilmList{_alias: alias}
	q.fid = path.NewInt16(q, "FID")
	q.title = path.NewString(q, "title")
	q.description = path.NewString(q, "description")
	q.category = path.NewString(q, "category")
	q.price = path.NewBinary(q, "price")
	q.length = path.NewInt16(q, "length")
	q.rating = path.NewString(q, "rating")
	q.actors = path.NewString(q, "actors")
	return q
}

type QNicerButSlowerFilmList struct {
	_alias      string
	fid         path.Int16
	title       path.String
	description path.String
	category    path.String
	price       path.Binary
	length      path.Int16
	rating      path.String
	actors      path.String
}

// core.Table Functions

func (q QNicerButSlowerFilmList) GetColumns() []core.Column {
	return []core.Column{
		q.fid,
		q.title,
		q.description,
		q.category,
		q.price,
		q.length,
		q.rating,
		q.actors,
	}
}

func (q QNicerButSlowerFilmList) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QNicerButSlowerFilmList) GetAlias() string {
	return q._alias
}

func (q QNicerButSlowerFilmList) GetName() string {
	return "nicer_but_slower_film_list"
}

func (q QNicerButSlowerFilmList) GetParent() string {
	return "sakila"
}

// Column Functions

func (q QNicerButSlowerFilmList) Fid() path.Int16 {
	return q.fid
}

func (q QNicerButSlowerFilmList) Title() path.String {
	return q.title
}

func (q QNicerButSlowerFilmList) Description() path.String {
	return q.description
}

func (q QNicerButSlowerFilmList) Category() path.String {
	return q.category
}

func (q QNicerButSlowerFilmList) Price() path.Binary {
	return q.price
}

func (q QNicerButSlowerFilmList) Length() path.Int16 {
	return q.length
}

func (q QNicerButSlowerFilmList) Rating() path.String {
	return q.rating
}

func (q QNicerButSlowerFilmList) Actors() path.String {
	return q.actors
}
