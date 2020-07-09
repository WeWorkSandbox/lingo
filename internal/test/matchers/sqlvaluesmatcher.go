package matchers

import (
	"fmt"
	"github.com/onsi/gomega"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"

	"github.com/weworksandbox/lingo/sql"
)

// SQLValuesMatcher matches a `sql.Data` `Values()` value
type SQLValuesMatcher struct {
	Expected interface{}
}

func (matcher *SQLValuesMatcher) Match(actual interface{}) (success bool, err error) {
	if isNil(actual) {
		return false, fmt.Errorf("expected a sql.Data, got nil")
	}

	s, ok := actual.(sql.Data)
	if !ok {
		return false, fmt.Errorf("expected a sql.Data.  Got:\n%s", format.Object(actual, 1))
	}

	var subMatcher types.GomegaMatcher
	var hasSubMatcher bool
	if matcher.Expected != nil {
		subMatcher, hasSubMatcher = (matcher.Expected).(types.GomegaMatcher)
		if !hasSubMatcher {
			subMatcher = gomega.Equal(matcher.Expected)
		}
		return subMatcher.Match(s.Values())
	}

	return false, fmt.Errorf(
		"SQLValuesMatcher must be passed zero or more multiple matchers.  Got:\n%s",
		format.Object(matcher.Expected, 1))
}

func (matcher *SQLValuesMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to match sql.Data values", matcher.Expected)
}

func (matcher *SQLValuesMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to match sql.Data values", matcher.Expected)
}
