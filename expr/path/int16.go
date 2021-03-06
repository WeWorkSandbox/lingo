// Code generated by an internal Lingo tool, genpaths.go - DO NOT EDIT

package path

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/expr/set"
	"github.com/weworksandbox/lingo/sql"
)

func NewInt16WithAlias(e lingo.Table, name, alias string) Int16 {
	return Int16{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewInt16(e lingo.Table, name string) Int16 {
	return NewInt16WithAlias(e, name, "")
}

type Int16 struct {
	entity lingo.Table
	name   string
	alias  string
}

func (p Int16) GetParent() lingo.Table {
	return p.entity
}

func (p Int16) GetName() string {
	return p.name
}

func (p Int16) GetAlias() string {
	return p.alias
}

func (p Int16) As(alias string) Int16 {
	p.alias = alias
	return p
}

func (p Int16) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, p)
}

func (p Int16) To(value int16) set.Set {
	return set.To(p, expr.NewValue(value))
}

func (p Int16) ToExpr(exp lingo.Expression) set.Set {
	return set.To(p, exp)
}

func (p Int16) Eq(value int16) operator.Binary {
	return operator.Eq(p, expr.NewValue(value))
}

func (p Int16) EqPath(exp lingo.Expression) operator.Binary {
	return operator.Eq(p, exp)
}

func (p Int16) NotEq(value int16) operator.Binary {
	return operator.NotEq(p, expr.NewValue(value))
}

func (p Int16) NotEqPath(exp lingo.Expression) operator.Binary {
	return operator.NotEq(p, exp)
}

func (p Int16) LT(value int16) operator.Binary {
	return operator.LessThan(p, expr.NewValue(value))
}

func (p Int16) LTPath(exp lingo.Expression) operator.Binary {
	return operator.LessThan(p, exp)
}

func (p Int16) LTOrEq(value int16) operator.Binary {
	return operator.LessThanOrEqual(p, expr.NewValue(value))
}

func (p Int16) LTOrEqPath(exp lingo.Expression) operator.Binary {
	return operator.LessThanOrEqual(p, exp)
}

func (p Int16) GT(value int16) operator.Binary {
	return operator.GreaterThan(p, expr.NewValue(value))
}

func (p Int16) GTPath(exp lingo.Expression) operator.Binary {
	return operator.GreaterThan(p, exp)
}

func (p Int16) GTOrEq(value int16) operator.Binary {
	return operator.GreaterThanOrEqual(p, expr.NewValue(value))
}

func (p Int16) GTOrEqPath(exp lingo.Expression) operator.Binary {
	return operator.GreaterThanOrEqual(p, exp)
}

func (p Int16) IsNull() operator.Unary {
	return operator.IsNull(p)
}

func (p Int16) IsNotNull() operator.Unary {
	return operator.IsNotNull(p)
}

func (p Int16) In(values ...int16) operator.Binary {
	return operator.In(p, expr.NewParens(expr.NewValue(values)))
}

func (p Int16) InPaths(exps ...lingo.Expression) operator.Binary {
	return operator.In(p, expr.NewParens(expr.ToList(exps)))
}

func (p Int16) NotIn(values ...int16) operator.Binary {
	return operator.NotIn(p, expr.NewParens(expr.NewValue(values)))
}

func (p Int16) NotInPaths(exps ...lingo.Expression) operator.Binary {
	return operator.NotIn(p, expr.NewParens(expr.ToList(exps)))
}

func (p Int16) Between(first, second int16) operator.Binary {
	return operator.Between(p, expr.NewValue(first), expr.NewValue(second))
}

func (p Int16) BetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.Between(p, first, second)
}

func (p Int16) NotBetween(first, second int16) operator.Binary {
	return operator.NotBetween(p, expr.NewValue(first), expr.NewValue(second))
}

func (p Int16) NotBetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.NotBetween(p, first, second)
}
