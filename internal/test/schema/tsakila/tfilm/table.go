// Code generated by Lingo for table sakila.film - DO NOT EDIT

// +build !nolingo

package tfilm

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/path"
	"github.com/weworksandbox/lingo/sql"
)

func As(alias string) TFilm {
	return newTFilm(alias)
}

func New() TFilm {
	return newTFilm("")
}

func newTFilm(alias string) TFilm {
	t := TFilm{
		_alias: alias,
	}
	t.filmId = path.NewInt16(t, "film_id")
	t.title = path.NewString(t, "title")
	t.description = path.NewString(t, "description")
	t.releaseYear = path.NewUnsupported(t, "release_year")
	t.languageId = path.NewInt8(t, "language_id")
	t.originalLanguageId = path.NewInt8(t, "original_language_id")
	t.rentalDuration = path.NewInt8(t, "rental_duration")
	t.rentalRate = path.NewBinary(t, "rental_rate")
	t.length = path.NewInt16(t, "length")
	t.replacementCost = path.NewBinary(t, "replacement_cost")
	t.rating = path.NewString(t, "rating")
	t.specialFeatures = path.NewString(t, "special_features")
	t.lastUpdate = path.NewTime(t, "last_update")
	return t
}

type TFilm struct {
	_alias string

	filmId             path.Int16
	title              path.String
	description        path.String
	releaseYear        path.Unsupported
	languageId         path.Int8
	originalLanguageId path.Int8
	rentalDuration     path.Int8
	rentalRate         path.Binary
	length             path.Int16
	replacementCost    path.Binary
	rating             path.String
	specialFeatures    path.String
	lastUpdate         path.Time
}

// lingo.Table Functions

func (t TFilm) GetColumns() []lingo.Column {
	return []lingo.Column{
		t.filmId,
		t.title,
		t.description,
		t.releaseYear,
		t.languageId,
		t.originalLanguageId,
		t.rentalDuration,
		t.rentalRate,
		t.length,
		t.replacementCost,
		t.rating,
		t.specialFeatures,
		t.lastUpdate,
	}
}

func (t TFilm) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, t)
}

func (t TFilm) GetAlias() string {
	return t._alias
}

func (t TFilm) GetName() string {
	return "film"
}

func (t TFilm) GetParent() string {
	return "sakila"
}

// Column Functions

func (t TFilm) FilmId() path.Int16 {
	return t.filmId
}

func (t TFilm) Title() path.String {
	return t.title
}

func (t TFilm) Description() path.String {
	return t.description
}

func (t TFilm) ReleaseYear() path.Unsupported {
	return t.releaseYear
}

func (t TFilm) LanguageId() path.Int8 {
	return t.languageId
}

func (t TFilm) OriginalLanguageId() path.Int8 {
	return t.originalLanguageId
}

func (t TFilm) RentalDuration() path.Int8 {
	return t.rentalDuration
}

func (t TFilm) RentalRate() path.Binary {
	return t.rentalRate
}

func (t TFilm) Length() path.Int16 {
	return t.length
}

func (t TFilm) ReplacementCost() path.Binary {
	return t.replacementCost
}

func (t TFilm) Rating() path.String {
	return t.rating
}

func (t TFilm) SpecialFeatures() path.String {
	return t.specialFeatures
}

func (t TFilm) LastUpdate() path.Time {
	return t.lastUpdate
}
