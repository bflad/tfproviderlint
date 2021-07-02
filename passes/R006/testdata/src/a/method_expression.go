package a

import (
	"a/methodexpression"
)

func fmethodexpression() *methodexpression.RetryError { // want "RetryFunc should include RetryableError\\(\\) handling or be removed"
	return methodexpression.RetryableError(nil)
}
