// Code generated by Lingo for table information_schema.OPTIMIZER_TRACE - DO NOT EDIT

package qoptimizertrace

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QOptimizerTrace {
	return instance
}

func Query() path.String {
	return instance.query
}

func Trace() path.String {
	return instance.trace
}

func MissingBytesBeyondMaxMemSize() path.Int {
	return instance.missingBytesBeyondMaxMemSize
}

func InsufficientPrivileges() path.Bool {
	return instance.insufficientPrivileges
}
