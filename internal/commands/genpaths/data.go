package main

import (
	"github.com/weworksandbox/lingo/internal/generator"
)

var pathData = []generator.Path{
	{
		Name:     "Binary",
		Filename: "binary.go",
		GoType:   "[]byte",
		Imports: []string{
			generator.PkgLingo,
			generator.PkgExp,
			generator.PkgOperator,
			generator.PkgSet,
			generator.PkgSQL,
		},
		Operations: generator.Ops{
			Set:        true,
			Equality:   true,
			Comparison: true,
			Nullable:   true,
			In:         true,
			Between:    true,
		},
	},
	{
		Name:     "Bool",
		Filename: "bool.go",
		GoType:   "bool",
		Imports: []string{
			generator.PkgLingo,
			generator.PkgExp,
			generator.PkgOperator,
			generator.PkgSet,
			generator.PkgSQL,
		},
		Operations: generator.Ops{
			Set:      true,
			Equality: true,
			Nullable: true,
		},
	},
	{
		Name:     "Float32",
		Filename: "float32.go",
		GoType:   "float32",
		Imports: []string{
			generator.PkgLingo,
			generator.PkgExp,
			generator.PkgOperator,
			generator.PkgSet,
			generator.PkgSQL,
		},
		Operations: generator.Ops{
			Set:        true,
			Equality:   true,
			Comparison: true,
			Nullable:   true,
			In:         true,
			Between:    true,
		},
	},
	{
		Name:     "Float64",
		Filename: "float64.go",
		GoType:   "float64",
		Imports: []string{
			generator.PkgLingo,
			generator.PkgExp,
			generator.PkgOperator,
			generator.PkgSet,
			generator.PkgSQL,
		},
		Operations: generator.Ops{
			Set:        true,
			Equality:   true,
			Comparison: true,
			Nullable:   true,
			In:         true,
			Between:    true,
		},
	},
	{
		Name:     "Int",
		Filename: "int.go",
		GoType:   "int",
		Imports: []string{
			generator.PkgLingo,
			generator.PkgExp,
			generator.PkgOperator,
			generator.PkgSet,
			generator.PkgSQL,
		},
		Operations: generator.Ops{
			Set:        true,
			Equality:   true,
			Comparison: true,
			Nullable:   true,
			In:         true,
			Between:    true,
		},
	},
	{
		Name:     "Int8",
		Filename: "int8.go",
		GoType:   "int8",
		Imports: []string{
			generator.PkgLingo,
			generator.PkgExp,
			generator.PkgOperator,
			generator.PkgSet,
			generator.PkgSQL,
		},
		Operations: generator.Ops{
			Set:        true,
			Equality:   true,
			Comparison: true,
			Nullable:   true,
			In:         true,
			Between:    true,
		},
	},
	{
		Name:     "Int16",
		Filename: "int16.go",
		GoType:   "int16",
		Imports: []string{
			generator.PkgLingo,
			generator.PkgExp,
			generator.PkgOperator,
			generator.PkgSet,
			generator.PkgSQL,
		},
		Operations: generator.Ops{
			Set:        true,
			Equality:   true,
			Comparison: true,
			Nullable:   true,
			In:         true,
			Between:    true,
		},
	},
	{
		Name:     "Int32",
		Filename: "int32.go",
		GoType:   "int32",
		Imports: []string{
			generator.PkgLingo,
			generator.PkgExp,
			generator.PkgOperator,
			generator.PkgSet,
			generator.PkgSQL,
		},
		Operations: generator.Ops{
			Set:        true,
			Equality:   true,
			Comparison: true,
			Nullable:   true,
			In:         true,
			Between:    true,
		},
	},
	{
		Name:     "Int64",
		Filename: "int64.go",
		GoType:   "int64",
		Imports: []string{
			generator.PkgLingo,
			generator.PkgExp,
			generator.PkgOperator,
			generator.PkgSet,
			generator.PkgSQL,
		},
		Operations: generator.Ops{
			Set:        true,
			Equality:   true,
			Comparison: true,
			Nullable:   true,
			In:         true,
			Between:    true,
		},
	},
	{
		Name:     "String",
		Filename: "string.go",
		GoType:   "string",
		Imports: []string{
			generator.PkgLingo,
			generator.PkgExp,
			generator.PkgOperator,
			generator.PkgSet,
			generator.PkgSQL,
		},
		Operations: generator.Ops{
			Set:        true,
			Equality:   true,
			Comparison: true,
			Nullable:   true,
			In:         true,
			Between:    true,
		},
	},
	{
		Name:     "Time",
		Filename: "time.go",
		GoType:   "time.Time",
		Imports: []string{
			"time",
			"",
			generator.PkgLingo,
			generator.PkgExp,
			generator.PkgOperator,
			generator.PkgSet,
			generator.PkgSQL,
		},
		Operations: generator.Ops{
			Set:        true,
			Equality:   true,
			Comparison: true,
			Nullable:   true,
			In:         true,
			Between:    true,
		},
	},
}