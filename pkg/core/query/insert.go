package query

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func InsertInto(entity core.Table) *InsertQuery {
	insert := InsertQuery{}
	insert.table = entity
	return &insert
}

type InsertQuery struct {
	table      core.Table
	columns    []core.Expression
	values     []core.Expression
	selectPart core.Expression
}

func (i *InsertQuery) Columns(columns ...core.Column) *InsertQuery {
	// TODO - validate we actually want to do this with our insert columns...
	i.columns = append(i.columns, convertToStringColumns(columns)...)
	return i
}

// ValuesConstants allows inserting constant values:
//
// INSERT INTO table1 (id, name, internal_name)
// values (123456, 'name1', 'internal_name');
func (i *InsertQuery) ValuesConstants(values ...interface{}) *InsertQuery {
	for _, value := range values {
		i.values = append(i.values, expression.NewValue(value))
	}
	return i
}

// Values allows inserting expressions with SQL functions:
//
// INSERT INTO table1 (uuid, name)
// values (UNHEX("1234567891234"), 'name1');
func (i *InsertQuery) Values(values ...core.Expression) *InsertQuery {
	i.values = append(i.values, values...)
	return i
}

// Select allows passing in a Select Query the following insert statements:
//
// INSERT INTO table1 (uuid, name)
// SELECT UNHEX("1234567891230"), a.name
// FROM table2 as a
// LEFT JOIN table1 as b ON a.name=b.remote_name
// WHERE b.remote_name = 'other_table';
func (i *InsertQuery) Select(q *SelectQuery) *InsertQuery {
	i.selectPart = q
	return i
}

func (i InsertQuery) ToSQL(d core.Dialect) (sql.Data, error) {
	var s = sql.String("INSERT INTO")

	if check.IsValueNilOrBlank(i.table) {
		return nil, expression.ErrorAroundSQL(expression.ExpressionIsNil("table"), s.String())
	}
	if i.table.GetAlias() != "" {
		return nil, expression.ErrorAroundSQL(errors.New("table alias must be unset"), s.String())
	}
	tableSQL, err := i.table.ToSQL(d)
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, s.String())
	}
	s = s.AppendWithSpace(tableSQL)

	if check.IsValueNilOrEmpty(i.columns) {
		return nil, expression.ErrorAroundSQL(expression.ExpressionCannotBeEmpty("columns"), s.String())
	}
	pathsSQL, err := JoinToSQL(d, sepPathComma, i.columns)
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, s.String())
	}
	s = s.SurroundAppend(" (", ")", pathsSQL) // Include space before first paren!

	if i.selectPart != nil {
		s, err = i.buildSelectFrom(d, s)
		if err != nil {
			return nil, err
		}
	} else {
		s, err = i.buildValues(d, s)
		if err != nil {
			return s, err
		}
	}

	return s, nil
}

func (i InsertQuery) buildSelectFrom(d core.Dialect, s sql.Data) (sql.Data, error) {
	selectSQL, err := i.selectPart.ToSQL(d)
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, s.String())
	}
	if selectSQL.String() != "" {
		s = s.AppendWithSpace(selectSQL)
	}
	return s, nil
}

func (i InsertQuery) buildValues(d core.Dialect, s sql.Data) (sql.Data, error) {
	colsLen := len(i.columns)
	valuesLen := len(i.values)
	if colsLen != valuesLen {
		err := fmt.Errorf("column count %d does not match values count %d", colsLen, valuesLen)
		return nil, expression.ErrorAroundSQL(err, s.String())
	}

	valuesSQL, err := JoinToSQL(d, sepPathComma, i.values)
	if err != nil {
		return nil, expression.ErrorAroundSQL(err, s.String())
	}
	if valuesSQL.String() != "" {
		s = s.AppendWithSpace(sql.String("VALUES")).SurroundAppend(" (", ")", valuesSQL)
	}
	return s, nil
}
