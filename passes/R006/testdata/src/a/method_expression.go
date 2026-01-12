package a

import (
	"testdata/src/a/methodexpression"
)

// With Go 1.25+/tools v0.40.0+, the analyzer correctly identifies
// methodexpression.RetryableError as resource.RetryableError through
// variable assignment tracking, so no diagnostic is expected.
func fmethodexpression() *methodexpression.RetryError {
	return methodexpression.RetryableError(nil)
}
