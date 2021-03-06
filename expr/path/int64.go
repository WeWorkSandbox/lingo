// Code generated by an internal Lingo tool, genpaths.go - DO NOT EDIT

package path

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/expr/set"
	"github.com/weworksandbox/lingo/sql"
)

func NewInt64WithAlias(e lingo.Table, name, alias string) Int64 {
	return Int64{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewInt64(e lingo.Table, name string) Int64 {
	return NewInt64WithAlias(e, name, "")
}

type Int64 struct {
	entity lingo.Table
	name   string
	alias  string
}

func (p Int64) GetParent() lingo.Table {
	return p.entity
}

func (p Int64) GetName() string {
	return p.name
}

func (p Int64) GetAlias() string {
	return p.alias
}

func (p Int64) As(alias string) Int64 {
	p.alias = alias
	return p
}

func (p Int64) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, p)
}

func (p Int64) To(value int64) set.Set {
	return set.To(p, expr.NewValue(value))
}

func (p Int64) ToExpr(exp lingo.Expression) set.Set {
	return set.To(p, exp)
}

func (p Int64) Eq(value int64) operator.Binary {
	return operator.Eq(p, expr.NewValue(value))
}

func (p Int64) EqPath(exp lingo.Expression) operator.Binary {
	return operator.Eq(p, exp)
}

func (p Int64) NotEq(value int64) operator.Binary {
	return operator.NotEq(p, expr.NewValue(value))
}

func (p Int64) NotEqPath(exp lingo.Expression) operator.Binary {
	return operator.NotEq(p, exp)
}

func (p Int64) LT(value int64) operator.Binary {
	return operator.LessThan(p, expr.NewValue(value))
}

func (p Int64) LTPath(exp lingo.Expression) operator.Binary {
	return operator.LessThan(p, exp)
}

func (p Int64) LTOrEq(value int64) operator.Binary {
	return operator.LessThanOrEqual(p, expr.NewValue(value))
}

func (p Int64) LTOrEqPath(exp lingo.Expression) operator.Binary {
	return operator.LessThanOrEqual(p, exp)
}

func (p Int64) GT(value int64) operator.Binary {
	return operator.GreaterThan(p, expr.NewValue(value))
}

func (p Int64) GTPath(exp lingo.Expression) operator.Binary {
	return operator.GreaterThan(p, exp)
}

func (p Int64) GTOrEq(value int64) operator.Binary {
	return operator.GreaterThanOrEqual(p, expr.NewValue(value))
}

func (p Int64) GTOrEqPath(exp lingo.Expression) operator.Binary {
	return operator.GreaterThanOrEqual(p, exp)
}

func (p Int64) IsNull() operator.Unary {
	return operator.IsNull(p)
}

func (p Int64) IsNotNull() operator.Unary {
	return operator.IsNotNull(p)
}

func (p Int64) In(values ...int64) operator.Binary {
	return operator.In(p, expr.NewParens(expr.NewValue(values)))
}

func (p Int64) InPaths(exps ...lingo.Expression) operator.Binary {
	return operator.In(p, expr.NewParens(expr.ToList(exps)))
}

func (p Int64) NotIn(values ...int64) operator.Binary {
	return operator.NotIn(p, expr.NewParens(expr.NewValue(values)))
}

func (p Int64) NotInPaths(exps ...lingo.Expression) operator.Binary {
	return operator.NotIn(p, expr.NewParens(expr.ToList(exps)))
}

func (p Int64) Between(first, second int64) operator.Binary {
	return operator.Between(p, expr.NewValue(first), expr.NewValue(second))
}

func (p Int64) BetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.Between(p, first, second)
}

func (p Int64) NotBetween(first, second int64) operator.Binary {
	return operator.NotBetween(p, expr.NewValue(first), expr.NewValue(second))
}

func (p Int64) NotBetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.NotBetween(p, first, second)
}
