// Code generated by Lingo for table information_schema.INNODB_SYS_TABLESPACES - DO NOT EDIT

package qinnodbsystablespaces

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QInnodbSysTablespaces {
	return instance
}

func Space() path.Int {
	return instance.space
}

func Name() path.String {
	return instance.name
}

func Flag() path.Int {
	return instance.flag
}

func FileFormat() path.String {
	return instance.fileFormat
}

func RowFormat() path.String {
	return instance.rowFormat
}

func PageSize() path.Int {
	return instance.pageSize
}

func ZipPageSize() path.Int {
	return instance.zipPageSize
}

func SpaceType() path.String {
	return instance.spaceType
}

func FsBlockSize() path.Int {
	return instance.fsBlockSize
}

func FileSize() path.Int64 {
	return instance.fileSize
}

func AllocatedSize() path.Int64 {
	return instance.allocatedSize
}
