package path_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/path"
	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/sql"
)

var _ = Describe("Table", func() {

	Context("ExpandTableWithDialect", func() {

		var (
			d     lingo.Dialect
			table lingo.Table

			sql sql.Data
			err error
		)

		BeforeEach(func() {
			d = expandTableDialectSuccess{}
			table = NewMockTable()
		})

		JustBeforeEach(func() {
			sql, err = path.ExpandTableWithDialect(d, table)
		})

		It("Returns valid SQL", func() {
			Expect(sql).To(matchers.MatchSQLString("expand table sql"))
		})

		It("Returns no error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		Context("`Dialect` does not support `ExpandTableDialect`", func() {

			BeforeEach(func() {
				d = NewMockDialect()
				pegomock.When(d.GetName()).ThenReturn("mock")
			})

			It("Returns a nil SQL", func() {
				Expect(sql).To(BeNil())
			})

			It("Returns a Dialect not supported error", func() {
				Expect(err).To(MatchError(matchers.EqString("dialect '%s' does not support '%s'", "mock", "path.ExpandTableDialect")))
			})
		})

		Context("Dialect returns an error", func() {

			BeforeEach(func() {
				d = expandTableDialectFailure{}
			})

			It("Returns a nil SQL", func() {
				Expect(sql).To(BeNil())
			})

			It("Returns the `Dialect` `ExpandTableDialect` error", func() {
				Expect(err).To(MatchError("expand table error"))
			})
		})
	})
})

type expandTableDialectSuccess struct{}

func (expandTableDialectSuccess) GetName() string { return "expand table dialect" }
func (expandTableDialectSuccess) ExpandTable(_ lingo.Table) (sql.Data, error) {
	return sql.String("expand table sql"), nil
}

type expandTableDialectFailure struct{ expandTableDialectSuccess }

func (expandTableDialectFailure) ExpandTable(_ lingo.Table) (sql.Data, error) {
	return nil, errors.New("expand table error")
}
