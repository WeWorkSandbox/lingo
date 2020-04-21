// Code generated by Lingo for table sakila.film_category - DO NOT EDIT

package qfilmcategory

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QFilmCategory {
	return newQFilmCategory(alias)
}

func New() QFilmCategory {
	return newQFilmCategory("")
}

func newQFilmCategory(alias string) QFilmCategory {
	q := QFilmCategory{_alias: alias}
	q.filmId = path.NewInt16Path(q, "film_id")
	q.categoryId = path.NewInt8Path(q, "category_id")
	q.lastUpdate = path.NewTimePath(q, "last_update")
	return q
}

type QFilmCategory struct {
	_alias     string
	filmId     path.Int16Path
	categoryId path.Int8Path
	lastUpdate path.TimePath
}

// core.Table Functions

func (q QFilmCategory) GetColumns() []core.Column {
	return []core.Column{
		q.filmId,
		q.categoryId,
		q.lastUpdate,
	}
}

func (q QFilmCategory) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QFilmCategory) GetAlias() string {
	return q._alias
}

func (q QFilmCategory) GetName() string {
	return "film_category"
}

func (q QFilmCategory) GetParent() string {
	return "sakila"
}

// Column Functions

func (q QFilmCategory) FilmId() path.Int16Path {
	return q.filmId
}

func (q QFilmCategory) CategoryId() path.Int8Path {
	return q.categoryId
}

func (q QFilmCategory) LastUpdate() path.TimePath {
	return q.lastUpdate
}
