// Code generated by Lingo for table information_schema.INNODB_SYS_FOREIGN_COLS - DO NOT EDIT

package qinnodbsysforeigncols

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QInnodbSysForeignCols {
	return newQInnodbSysForeignCols(alias)
}

func New() QInnodbSysForeignCols {
	return newQInnodbSysForeignCols("")
}

func newQInnodbSysForeignCols(alias string) QInnodbSysForeignCols {
	q := QInnodbSysForeignCols{_alias: alias}
	q.id = path.NewStringPath(q, "ID")
	q.forColName = path.NewStringPath(q, "FOR_COL_NAME")
	q.refColName = path.NewStringPath(q, "REF_COL_NAME")
	q.pos = path.NewIntPath(q, "POS")
	return q
}

type QInnodbSysForeignCols struct {
	_alias     string
	id         path.String
	forColName path.String
	refColName path.String
	pos        path.Int
}

// core.Table Functions

func (q QInnodbSysForeignCols) GetColumns() []core.Column {
	return []core.Column{
		q.id,
		q.forColName,
		q.refColName,
		q.pos,
	}
}

func (q QInnodbSysForeignCols) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QInnodbSysForeignCols) GetAlias() string {
	return q._alias
}

func (q QInnodbSysForeignCols) GetName() string {
	return "INNODB_SYS_FOREIGN_COLS"
}

func (q QInnodbSysForeignCols) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QInnodbSysForeignCols) Id() path.String {
	return q.id
}

func (q QInnodbSysForeignCols) ForColName() path.String {
	return q.forColName
}

func (q QInnodbSysForeignCols) RefColName() path.String {
	return q.refColName
}

func (q QInnodbSysForeignCols) Pos() path.Int {
	return q.pos
}
