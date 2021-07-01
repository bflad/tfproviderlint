package a

import (
	"a/methodexpression"
)

func fmethodexpression() *methodexpression.RetryError {
	return methodexpression.RetryableError(nil)
}
