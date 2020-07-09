package json_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expr/json"
	. "github.com/weworksandbox/lingo/pkg/core/expr/matchers"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

var _ = Describe("JSON", func() {

	Context("#NewJSONOperation", func() {

		var (
			left        core.Expression
			op          json.Operand
			expressions []core.Expression

			operation core.ComboExpression
		)

		BeforeEach(func() {
			left = NewMockExpression()
			pegomock.When(left.ToSQL(AnyCoreDialect())).ThenReturn(sql.String("left sql"), nil)

			op = json.Extract
			expressions = []core.Expression{
				NewMockExpression(),
				NewMockExpression(),
			}
			pegomock.When(expressions[0].ToSQL(AnyCoreDialect())).ThenReturn(sql.String("expressions[0]"), nil)
			pegomock.When(expressions[1].ToSQL(AnyCoreDialect())).ThenReturn(sql.String("expressions[1]"), nil)
		})

		JustBeforeEach(func() {
			operation = json.NewJSONOperation(left, op, expressions...)
		})

		It("Creates a valid operation", func() {
			Expect(operation).ToNot(BeNil())
		})

		Context("#ToSQL", func() {

			var (
				d core.Dialect

				sql sql.Data
				err error
			)

			BeforeEach(func() {
				d = jsonDialectSuccess{}
			})

			JustBeforeEach(func() {
				sql, err = operation.ToSQL(d)
			})

			It("Returns a valid SQL", func() {
				Expect(sql).To(matchers.MatchSQLString("json sql"))
			})

			It("Does not error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			Context("Dialect does not support `JSONOperator`", func() {

				BeforeEach(func() {
					d = NewMockDialect()
					pegomock.When(d.GetName()).ThenReturn("mock")
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(ContainSubstring("dialect '%s' does not support '%s'", "mock", "json.Dialect")))
				})
			})

			Context("left is nil", func() {

				BeforeEach(func() {
					left = nil
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns an error that left is nil", func() {
					Expect(err).To(MatchError("left of 'json' cannot be empty"))
				})
			})

			Context("left returns an error", func() {

				BeforeEach(func() {
					leftErr := errors.New("left error")
					pegomock.When(left.ToSQL(AnyCoreDialect())).ThenReturn(nil, leftErr)
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns the left error", func() {
					Expect(err).To(MatchError(ContainSubstring("left error")))
				})
			})

			Context("an expr is nil", func() {

				BeforeEach(func() {
					expressions[1] = nil
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns an error that an expr is nil", func() {
					Expect(err).To(MatchError("expressions[1] of json cannot be empty"))
				})
			})

			Context("an expr returns an error", func() {

				BeforeEach(func() {
					expErr := errors.New("exp error")
					pegomock.When(expressions[1].ToSQL(AnyCoreDialect())).ThenReturn(nil, expErr)
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns the expr error", func() {
					Expect(err).To(MatchError(ContainSubstring("exp error")))
				})
			})

			Context("JSONOperator returns an error", func() {

				BeforeEach(func() {
					d = jsonDialectFailure{}
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns JSONOperator error", func() {
					Expect(err).To(MatchError(ContainSubstring("json failure")))
				})
			})
		})
	})
})

type jsonDialectSuccess struct{}

func (jsonDialectSuccess) GetName() string { return "json success" }
func (jsonDialectSuccess) JSONOperator(sql.Data, json.Operand, []sql.Data) (sql.Data, error) {
	return sql.String("json sql"), nil
}

type jsonDialectFailure struct{ jsonDialectSuccess }

func (jsonDialectFailure) JSONOperator(sql.Data, json.Operand, []sql.Data) (sql.Data, error) {
	return nil, errors.New("json failure")
}