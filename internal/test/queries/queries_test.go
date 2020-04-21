package queries_test

import (
	"strings"
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"

	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/internal/test/runner"
	"github.com/weworksandbox/lingo/pkg/core"
)

// Query is used by Functional tests, along with benchmark tests. They are used for setting up common data to
// ensure performance and code quality.
type Query struct {
	Name      string
	Focus     bool
	Benchmark bool

	// Params used during the test
	Params
}

type Params struct {
	Dialect      core.Dialect
	SQL          func() core.Expression
	SQLAssert    types.GomegaMatcher
	ValuesAssert types.GomegaMatcher
	ErrAssert    types.GomegaMatcher
}

func BenchmarkQueries(b *testing.B) {
	b.ReportAllocs()

	for _, query := range allQueries {
		if !query.Benchmark {
			b.Skip("Benchmark turned off for query ", query.Name)
		}

		if query.Dialect == nil {
			b.Errorf("Query '%s' does not have a Dialect", query.Name)
		}

		b.Run(query.Name, func(parallel *testing.B) {
			parallel.ReportAllocs()
			parallel.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, _ = query.SQL().GetSQL(query.Dialect)
				}
			})
		})
	}
}

func TestQueries(t *testing.T) {
	var _ = ginkgo.Describe("Queries", func() {
		table.DescribeTable("query.go",
			func(p Params) {
				// Sanity check
				Expect(p).ToNot(BeNil())
				Expect(p.Dialect).ToNot(BeNil(), "Dialect was nil")
				Expect(p.SQL).ToNot(BeNil(), "SQL was nil")
				Expect(p.SQLAssert).ToNot(BeNil(), "SQLAssert was nil")
				Expect(p.ErrAssert).ToNot(BeNil(), "ErrAssert was nil for ")

				sql, err := p.SQL().GetSQL(p.Dialect)
				Expect(err).To(p.ErrAssert)
				Expect(sql).To(MatchSQLString(p.SQLAssert))
				Expect(sql).To(MatchSQLValues(p.ValuesAssert))
			},
			acceptanceEntries...,
		)
	})

	runner.SetupAndRunUnit(t, "Queries", "functional")
}

var (
	allQueries        = aggregateQueries(selectQueries)
	acceptanceEntries = queriesToEntries(allQueries)
)

func aggregateQueries(q ...[]Query) []Query {
	var result []Query
	for idx := range q {
		result = append(result, q[idx]...)
	}
	return result
}

func queriesToEntries(queries []Query) []table.TableEntry {
	var entries = make([]table.TableEntry, len(queries))
	for idx, query := range queries {
		entries[idx] = table.TableEntry{
			Description: query.Name,
			Parameters:  []interface{}{query.Params},
			Pending:     false,
			Focused:     query.Focus,
		}
	}
	return entries
}

// trimQuery replaces newlines with spaces, and removing any tabs. This way, SQL.SQL can use backticks.
func trimQuery(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", "")
	return strings.TrimSpace(s)
}
