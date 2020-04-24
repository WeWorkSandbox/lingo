// Code generated by Lingo for table information_schema.INNODB_SYS_FOREIGN - DO NOT EDIT

package qinnodbsysforeign

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QInnodbSysForeign {
	return newQInnodbSysForeign(alias)
}

func New() QInnodbSysForeign {
	return newQInnodbSysForeign("")
}

func newQInnodbSysForeign(alias string) QInnodbSysForeign {
	q := QInnodbSysForeign{_alias: alias}
	q.id = path.NewStringPath(q, "ID")
	q.forName = path.NewStringPath(q, "FOR_NAME")
	q.refName = path.NewStringPath(q, "REF_NAME")
	q.nCols = path.NewIntPath(q, "N_COLS")
	q.__type = path.NewIntPath(q, "TYPE")
	return q
}

type QInnodbSysForeign struct {
	_alias  string
	id      path.String
	forName path.String
	refName path.String
	nCols   path.Int
	__type  path.Int
}

// core.Table Functions

func (q QInnodbSysForeign) GetColumns() []core.Column {
	return []core.Column{
		q.id,
		q.forName,
		q.refName,
		q.nCols,
		q.__type,
	}
}

func (q QInnodbSysForeign) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QInnodbSysForeign) GetAlias() string {
	return q._alias
}

func (q QInnodbSysForeign) GetName() string {
	return "INNODB_SYS_FOREIGN"
}

func (q QInnodbSysForeign) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QInnodbSysForeign) Id() path.String {
	return q.id
}

func (q QInnodbSysForeign) ForName() path.String {
	return q.forName
}

func (q QInnodbSysForeign) RefName() path.String {
	return q.refName
}

func (q QInnodbSysForeign) NCols() path.Int {
	return q.nCols
}

func (q QInnodbSysForeign) Type() path.Int {
	return q.__type
}
