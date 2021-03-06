package query

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/check"
	"github.com/weworksandbox/lingo/expr/path"
	"github.com/weworksandbox/lingo/sql"
)

// NewNamedOnlyColumn creates a `lingo.Column` of which only the name of the column is filled out.
// Thus, when `ToSQL()` is called, only a single SQL with the value of `name` is returned.
func NewNamedOnlyColumn(name, parent string) lingo.Column {
	return &stringColumn{name: name, parent: stringParent{name: parent}}
}

type stringColumn struct {
	name   string
	parent stringParent
}

func (s stringColumn) ToSQL(_ lingo.Dialect) (sql.Data, error) {
	return sql.String(s.GetName()), nil
}
func (s stringColumn) GetName() string        { return s.name }
func (s stringColumn) GetParent() lingo.Table { return s.parent }
func (stringColumn) GetAlias() string         { return "" }

type stringParent struct {
	name string
}

func (s stringParent) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return path.ExpandTableWithDialect(d, s)
}

func (stringParent) GetAlias() string           { return "" }
func (s stringParent) GetName() string          { return s.name }
func (stringParent) GetColumns() []lingo.Column { return []lingo.Column{} }
func (stringParent) GetParent() string          { return "" }

func convertToStringColumns(columns []lingo.Column) []lingo.Expression {
	if check.IsValueNilOrBlank(columns) {
		return nil
	}

	var expressions = make([]lingo.Expression, 0, len(columns))
	for _, column := range columns {
		if check.IsValueNilOrBlank(column) {
			return nil
		}
		// TODO we might not even need this entire function or file... maybe remove?
		expressions = append(expressions, NewNamedOnlyColumn(column.GetName(), column.GetParent().GetName()))
	}
	return expressions
}
