package query_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/dialect"
	"github.com/weworksandbox/lingo/expr"
	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/query"
	"github.com/weworksandbox/lingo/query/matchers"
	"github.com/weworksandbox/lingo/sql"
)

var _ = Describe("Insert", func() {

	Context("InsertInto #Columns #Values", func() {

		var (
			table            lingo.Table
			cols             []lingo.Column
			valueExpressions []lingo.Expression
			valueConstants   []interface{}

			q *query.InsertQuery
		)

		BeforeEach(func() {
			table = NewMockTable()

			cols = []lingo.Column{
				NewMockColumn(),
				NewMockColumn(),
			}
			pegomock.When(cols[0].ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("cols[0].sqlStr"), nil)
			pegomock.When(cols[0].GetAlias()).ThenReturn("col[0].alias")
			pegomock.When(cols[0].GetName()).ThenReturn("col[0].name")
			pegomock.When(cols[0].GetParent()).ThenReturn(table)
			pegomock.When(cols[1].ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("cols[1].sqlStr"), nil)
			pegomock.When(cols[1].GetAlias()).ThenReturn("col[1].alias")
			pegomock.When(cols[1].GetName()).ThenReturn("col[1].name")
			pegomock.When(cols[1].GetParent()).ThenReturn(table)

			pegomock.When(table.GetColumns()).ThenReturn(cols)
			pegomock.When(table.ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("table.sqlStr"), nil)

			valueExpressions = []lingo.Expression{
				NewMockExpression(),
				NewMockExpression(),
			}
			pegomock.When(valueExpressions[0].ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("valueExpressions[0].sqlStr"), nil)
			pegomock.When(valueExpressions[1].ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("valueExpressions[1].sqlStr"), nil)

			// Ensure we reset valueConstants. Had random test failures due to this. Only happened on certain
			// random num test seeds.
			valueConstants = nil
		})

		JustBeforeEach(func() {
			q = query.InsertInto(table).Values(valueExpressions...).ValuesConstants(valueConstants...)
			if table != nil {
				q = q.Columns(table.GetColumns()...)
			}
		})

		Context("#ToSQL", func() {

			var (
				d lingo.Dialect

				sql sql.Data
				err error
			)

			BeforeEach(func() {
				d = dialect.Default{}
			})

			JustBeforeEach(func() {
				sql, err = q.ToSQL(d)
			})

			It("Returns a valid SQL string", func() {
				Expect(sql).To(MatchSQLString("INSERT INTO table.sqlStr (col[0].name, col[1].name) VALUES (valueExpressions[0].sqlStr, valueExpressions[1].sqlStr)"))
			})

			It("Returns no error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			Context("Table is nil", func() {

				BeforeEach(func() {
					table = nil
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns a table is nil error", func() {
					Expect(err).To(MatchError(ContainSubstring("table cannot be empty")))
				})
			})

			Context("Table has an alias", func() {

				BeforeEach(func() {
					pegomock.When(table.GetAlias()).ThenReturn("table alias")
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns a table alias cannot be set error", func() {
					Expect(err).To(MatchError(ContainSubstring("table alias must be unset")))
				})
			})

			Context("Table ToSQL has an error", func() {

				BeforeEach(func() {
					pegomock.When(table.ToSQL(matchers.AnyLingoDialect())).ThenReturn(nil, errors.New("table error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns the table error", func() {
					Expect(err).To(MatchError(ContainSubstring("table error")))
				})
			})

			Context("Columns are nil", func() {

				BeforeEach(func() {
					pegomock.When(table.GetColumns()).ThenReturn(nil)
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns a columns is nil error", func() {
					Expect(err).To(MatchError(ContainSubstring("expr '%s' cannot be empty", "columns")))
				})
			})

			// This will probably either be removed or need to change.
			// We are converting the columns passed in to StringColumn types.
			// Doing this, we dont have a 'clean' way to mock this out.
			XContext("Columns return an error", func() {

				BeforeEach(func() {
					pegomock.When(cols[len(cols)-1].ToSQL(matchers.AnyLingoDialect())).ThenReturn(nil, errors.New("col error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns the column error", func() {
					Expect(err).To(MatchError(ContainSubstring("col error")))
				})
			})

			Context("Values are nil", func() {

				BeforeEach(func() {
					valueExpressions = nil
				})

				It("Returns a nil SQL string", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns an error about column count vs values count", func() {
					Expect(err).To(MatchError(ContainSubstring("column count 2 does not match values count 0")))
				})
			})

			Context("Values return an error", func() {

				BeforeEach(func() {
					pegomock.When(valueExpressions[len(valueExpressions)-1].ToSQL(matchers.AnyLingoDialect())).ThenReturn(nil, errors.New("valueExpressions error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns the valueExpressions error", func() {
					Expect(err).To(MatchError(ContainSubstring("valueExpressions error")))
				})
			})

			Context("With value constants", func() {

				BeforeEach(func() {
					valueConstants = []interface{}{
						"stringHere",
						1.4e2,
					}
					valueExpressions = nil
				})

				It("Returns a valid SQL string", func() {
					Expect(sql).To(MatchSQLString("INSERT INTO table.sqlStr (col[0].name, col[1].name) VALUES (?, ?)"))
				})

				It("Returns valid SQL Values", func() {
					Expect(sql).To(MatchSQLValues(ConsistOf("stringHere", 1.4e2)))
				})

				It("Returns no error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})
		})

		Context("With select part", func() {

			var (
				sTable lingo.Table

				sq *query.SelectQuery
			)

			BeforeEach(func() {
				sTable = NewMockTable()
				pegomock.When(sTable.ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("select.sqlStr"), nil)

				sq = query.Select(expr.Star()).From(sTable)
			})

			JustBeforeEach(func() {
				q = q.Select(sq)
			})

			Context("#ToSQL", func() {

				var (
					d lingo.Dialect

					sql sql.Data
					err error
				)

				BeforeEach(func() {
					d = dialect.Default{}
				})

				JustBeforeEach(func() {
					sql, err = q.ToSQL(d)
				})

				It("Returns a valid SQL string", func() {
					Expect(sql).To(MatchSQLString("INSERT INTO table.sqlStr (col[0].name, col[1].name) SELECT * FROM select.sqlStr"))
				})

				It("Returns no error", func() {
					Expect(err).ToNot(HaveOccurred())
				})

				Context("Select returns an error", func() {

					BeforeEach(func() {
						pegomock.When(sTable.ToSQL(matchers.AnyLingoDialect())).ThenReturn(nil, errors.New("select error"))
					})

					It("Returns a nil SQL", func() {
						Expect(sql).To(BeNil())
					})

					It("Returns the select error", func() {
						Expect(err).To(MatchError(ContainSubstring("select error")))
					})
				})
			})
		})
	})
})
