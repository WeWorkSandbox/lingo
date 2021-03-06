// Code generated by Lingo for table sakila.language - DO NOT EDIT

// +build !nolingo

package tlanguage

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/path"
	"github.com/weworksandbox/lingo/sql"
)

func As(alias string) TLanguage {
	return newTLanguage(alias)
}

func New() TLanguage {
	return newTLanguage("")
}

func newTLanguage(alias string) TLanguage {
	t := TLanguage{
		_alias: alias,
	}
	t.languageId = path.NewInt8(t, "language_id")
	t.name = path.NewString(t, "name")
	t.lastUpdate = path.NewTime(t, "last_update")
	return t
}

type TLanguage struct {
	_alias string

	languageId path.Int8
	name       path.String
	lastUpdate path.Time
}

// lingo.Table Functions

func (t TLanguage) GetColumns() []lingo.Column {
	return []lingo.Column{
		t.languageId,
		t.name,
		t.lastUpdate,
	}
}

func (t TLanguage) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, t)
}

func (t TLanguage) GetAlias() string {
	return t._alias
}

func (t TLanguage) GetName() string {
	return "language"
}

func (t TLanguage) GetParent() string {
	return "sakila"
}

// Column Functions

func (t TLanguage) LanguageId() path.Int8 {
	return t.languageId
}

func (t TLanguage) Name() path.String {
	return t.name
}

func (t TLanguage) LastUpdate() path.Time {
	return t.lastUpdate
}
