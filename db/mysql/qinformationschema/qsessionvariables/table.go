// Code generated by Lingo for table information_schema.SESSION_VARIABLES - DO NOT EDIT

package qsessionvariables

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QSessionVariables {
	return newQSessionVariables(alias)
}

func New() QSessionVariables {
	return newQSessionVariables("")
}

func newQSessionVariables(alias string) QSessionVariables {
	q := QSessionVariables{_alias: alias}
	q.variableName = path.NewStringPath(q, "VARIABLE_NAME")
	q.variableValue = path.NewStringPath(q, "VARIABLE_VALUE")
	return q
}

type QSessionVariables struct {
	_alias        string
	variableName  path.StringPath
	variableValue path.StringPath
}

// core.Table Functions

func (q QSessionVariables) GetColumns() []core.Column {
	return []core.Column{
		q.variableName,
		q.variableValue,
	}
}

func (q QSessionVariables) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QSessionVariables) GetAlias() string {
	return q._alias
}

func (q QSessionVariables) GetName() string {
	return "SESSION_VARIABLES"
}

func (q QSessionVariables) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QSessionVariables) VariableName() path.StringPath {
	return q.variableName
}

func (q QSessionVariables) VariableValue() path.StringPath {
	return q.variableValue
}