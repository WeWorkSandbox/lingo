// Code generated by Lingo for table information_schema.INNODB_CMP_PER_INDEX - DO NOT EDIT

package qinnodbcmpperindex

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QInnodbCmpPerIndex {
	return instance
}

func DatabaseName() path.StringPath {
	return instance.databaseName
}

func TableName() path.StringPath {
	return instance.tableName
}

func IndexName() path.StringPath {
	return instance.indexName
}

func CompressOps() path.IntPath {
	return instance.compressOps
}

func CompressOpsOk() path.IntPath {
	return instance.compressOpsOk
}

func CompressTime() path.IntPath {
	return instance.compressTime
}

func UncompressOps() path.IntPath {
	return instance.uncompressOps
}

func UncompressTime() path.IntPath {
	return instance.uncompressTime
}