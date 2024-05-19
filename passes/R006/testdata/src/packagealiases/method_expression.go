package a

import (
	"testdata/src/a/methodexpression"
)

func fmethodexpression() *methodexpression.RetryError {
	return methodexpression.RetryableError(nil)
}
