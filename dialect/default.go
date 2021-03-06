package dialect

import (
	"fmt"
	"strings"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/join"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/expr/sort"
	"github.com/weworksandbox/lingo/query"
	"github.com/weworksandbox/lingo/sql"
)

// NewDefault takes options to configure a Default schema
func NewDefault(opts ...Option) (Default, error) {
	var o options
	for idx := range opts {
		if err := opts[idx](&o); err != nil {
			return Default{}, fmt.Errorf("unable to create default dialect: %w", err)
		}
	}
	return Default{
		includeSchemaName: o.includeSchemaName,
	}, nil
}

// Default schema uses the generic schema methods to work as a basic ANSI schema.
type Default struct {
	includeSchemaName bool
}

func (Default) GetName() string {
	return "Default"
}

func (Default) ValueFormat(count int) sql.Data {
	if count == 0 {
		return sql.Empty()
	}

	const (
		qMark = "?"
		comSp = ", " + qMark
	)

	var s strings.Builder

	numCommas := (count - 1) * len(comSp) // Subtract 1 cuz we add the len of the first question mark next
	s.Grow(numCommas + len(qMark))        // Add the first question mark

	_, _ = s.WriteString(qMark)
	for idx := 1; idx < count; idx++ {
		_, _ = s.WriteString(comSp)
	}
	return sql.String(s.String())
}

func (Default) SetValueFormat() string {
	return "="
}

func (d Default) ExpandTable(table lingo.Table) (sql.Data, error) {
	if d.includeSchemaName {
		return ExpandTableWithSchema(table)
	}
	return ExpandTable(table)
}

func (Default) ExpandColumn(column lingo.Column) (sql.Data, error) {
	return ExpandColumnWithParent(column)
}

func (Default) UnaryOperator(left sql.Data, op operator.Operator) (sql.Data, error) {
	return UnaryOperator(left, op)
}

func (Default) BinaryOperator(left sql.Data, op operator.Operator, right sql.Data) (sql.Data, error) {
	return BinaryOperator(left, op, right)
}

func (Default) VariadicOperator(left sql.Data, op operator.Operator, values []sql.Data) (sql.Data, error) {
	return VariadicOperator(left, op, values)
}

func (d Default) Value(value []interface{}) (sql.Data, error) {
	return Value(d, value)
}

func (Default) Join(left sql.Data, joinType join.Type, on sql.Data) (sql.Data, error) {
	return Join(left, joinType, on)
}

func (Default) OrderBy(left sql.Data, direction sort.Direction) (sql.Data, error) {
	return OrderBy(left, direction)
}

func (d Default) Set(left sql.Data, value sql.Data) (sql.Data, error) {
	return Set(d, left, value)
}

// Modify will build: [LIMIT limit] [OFFSET offset]
func (d Default) Modify(m query.Modifier) (sql.Data, error) {
	limit, lWasSet := m.Limit()
	offset, oWasSet := m.Offset()

	s := sql.Empty()
	if lWasSet {
		limitSQL, err := d.Value([]interface{}{limit})
		if err != nil {
			return nil, err
		}
		s = sql.String("LIMIT").AppendWithSpace(limitSQL)
	}
	if oWasSet {
		offsetSQL, err := d.Value([]interface{}{offset})
		if err != nil {
			return nil, err
		}
		s = s.AppendWithSpace(sql.String("OFFSET").AppendWithSpace(offsetSQL))
	}
	return s, nil
}
