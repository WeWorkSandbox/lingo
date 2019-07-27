// Code generated by Lingo for table information_schema.INNODB_SYS_TABLES - DO NOT EDIT

package qinnodbsystables

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QInnodbSysTables {
	return instance
}

func TableId() path.Int64Path {
	return instance.tableId
}

func Name() path.StringPath {
	return instance.name
}

func Flag() path.IntPath {
	return instance.flag
}

func NCols() path.IntPath {
	return instance.nCols
}

func Space() path.IntPath {
	return instance.space
}

func FileFormat() path.StringPath {
	return instance.fileFormat
}

func RowFormat() path.StringPath {
	return instance.rowFormat
}

func ZipPageSize() path.IntPath {
	return instance.zipPageSize
}

func SpaceType() path.StringPath {
	return instance.spaceType
}
