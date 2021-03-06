package query_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/query"
	"github.com/weworksandbox/lingo/sql"
)

var _ = Describe("where", func() {

	Context("#BuildWhereSQL", func() {

		var (
			d      lingo.Dialect
			values []lingo.Expression

			s   sql.Data
			err error
		)

		BeforeEach(func() {
			d = whereDialectSuccess{}
			values = []lingo.Expression{
				NewMockExpression(),
				NewMockExpression(),
				NewMockExpression(),
			}
			pegomock.When(values[0].ToSQL(d)).ThenReturn(sql.String("where 0 sqlStr"), nil)
			pegomock.When(values[1].ToSQL(d)).ThenReturn(sql.String("where 1 sqlStr"), nil)
			pegomock.When(values[2].ToSQL(d)).ThenReturn(sql.String("where 2 sqlStr"), nil)
		})

		JustBeforeEach(func() {
			s, err = query.BuildWhereSQL(d, values)
		})

		It("Combines all SQL with commas and `WHERE`", func() {
			Expect(s).To(matchers.MatchSQLString("WHERE where 0 sqlStr op where 1 sqlStr op where 2 sqlStr"))
		})

		It("Returns no error", func() {
			Expect(err).ShouldNot(HaveOccurred())
		})

		Context("With an error returning", func() {

			BeforeEach(func() {
				pegomock.When(values[2].ToSQL(d)).ThenReturn(nil, errors.New("last error"))
			})

			It("Returns a nil SQL", func() {
				Expect(s).To(BeNil())
			})

			It("Returns the error", func() {
				Expect(err).To(MatchError(ContainSubstring("last error")))
			})
		})

		Context("With 2 values", func() {

			BeforeEach(func() {
				values = values[:len(values)-1]
			})

			It("Combines all SQL with commas and `WHERE`", func() {
				Expect(s).To(matchers.MatchSQLString("WHERE where 0 sqlStr op where 1 sqlStr"))
			})

			It("Returns no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("With 1 value", func() {

			BeforeEach(func() {
				values = values[:1]
			})

			It("Combines all SQL with commas and `WHERE`", func() {
				Expect(s).To(matchers.MatchSQLString("WHERE where 0 sqlStr"))
			})

			It("Returns no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})

			Context("With an error returning", func() {

				BeforeEach(func() {
					pegomock.When(values[0].ToSQL(d)).ThenReturn(nil, errors.New("last error"))
				})

				It("Returns a nil SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns the error", func() {
					Expect(err).To(MatchError(ContainSubstring("last error")))
				})
			})
		})

		Context("With 0 values", func() {

			BeforeEach(func() {
				values = []lingo.Expression{}
			})

			It("Combines all SQL with commas and `WHERE`", func() {
				Expect(s).To(matchers.MatchSQLString(""))
			})

			It("Returns no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})
})

type whereDialectSuccess struct{}

func (whereDialectSuccess) GetName() string { return "where dialect success" }

func (s whereDialectSuccess) UnaryOperator(left sql.Data, _ operator.Operator) (sql.Data, error) {
	return left, nil
}

func (s whereDialectSuccess) BinaryOperator(left sql.Data, _ operator.Operator, value sql.Data) (sql.Data, error) {
	return left.AppendWithSpace(sql.String("op")).AppendWithSpace(value), nil
}

func (s whereDialectSuccess) VariadicOperator(left sql.Data, _ operator.Operator, values []sql.Data) (sql.Data, error) {
	var result = left
	for _, value := range values {
		result = result.AppendWithSpace(sql.String("op")).AppendWithSpace(value)
	}
	return result, nil
}
