package query_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo"
	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/query"
	"github.com/weworksandbox/lingo/sql"
)

var _ = Describe("Paths", func() {

	Context("ExpandTables", func() {

		var (
			paths []lingo.Expression

			result []lingo.Expression
		)

		BeforeEach(func() {
			col1 := NewMockColumn()
			col2 := NewMockColumn()
			table := NewMockTable()
			paths = []lingo.Expression{
				table,
			}
			pegomock.When(table.GetColumns()).ThenReturn([]lingo.Column{col1, col2})
		})

		JustBeforeEach(func() {
			result = query.ExpandTables(paths)
		})

		It("Returns 2 column expressions", func() {
			Expect(result).To(HaveLen(2))
		})

		Context("With `lingo.Expression`", func() {

			BeforeEach(func() {
				paths = []lingo.Expression{
					NewMockColumn(),
					NewMockExpression(),
				}
			})

			It("Returns two expressions", func() {
				Expect(result).To(HaveLen(2))
				ExpectWithOffset(0, result).To(ContainElement(paths[0]))
				ExpectWithOffset(1, result).To(ContainElement(paths[1]))
			})
		})
	})

	Context("#JoinToSQL", func() {

		var (
			d     lingo.Dialect
			sep   string
			paths []lingo.Expression

			s   sql.Data
			err error
		)

		BeforeEach(func() {
			d = NewMockDialect()
			sep = ".SEP."
			paths = []lingo.Expression{
				NewMockExpression(),
				NewMockExpression(),
			}
			pegomock.When(paths[0].ToSQL(d)).ThenReturn(sql.String("exp 1 sqlStr"), nil)
			pegomock.When(paths[1].ToSQL(d)).ThenReturn(sql.String("exp 2 sqlStr"), nil)
		})

		JustBeforeEach(func() {
			s, err = query.JoinToSQL(d, sep, paths)
		})

		It("Returns a combined SQL", func() {
			Expect(s).To(MatchSQLString("exp 1 sqlStr.SEP.exp 2 sqlStr"))
		})

		It("Returns no error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		Context("With one column", func() {

			BeforeEach(func() {
				paths = paths[:1]
			})

			It("Returns the original SQL with no sep", func() {
				Expect(s).To(MatchSQLString("exp 1 sqlStr"))
			})

			It("Returns no error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("With nil columns", func() {

			BeforeEach(func() {
				paths = nil
			})

			It("Returns an empty SQL", func() {
				Expect(s).To(MatchSQLString(""))
			})

			It("Returns a nil error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("With an error on the second expr", func() {

			BeforeEach(func() {
				pegomock.When(paths[1].ToSQL(d)).ThenReturn(nil, errors.New("second exp error"))
			})

			It("Returns a nil SQL", func() {
				Expect(s).To(BeNil())
			})

			It("Returns the second error", func() {
				Expect(err).To(MatchError(ContainSubstring("second exp error")))
			})
		})
	})
})
