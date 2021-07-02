package a

import (
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

//lintignore:R006
func commentIgnore_v2() *resource.RetryError {
	return nil
}

func failingAnonymousRetryFunc_v2() {
	_ = resource.Retry(1*time.Minute, func() *resource.RetryError { // want "RetryFunc should include RetryableError\\(\\) handling or be removed"
		return nil
	})
}

func failingNoCallExpr_v2() *resource.RetryError { // want "RetryFunc should include RetryableError\\(\\) handling or be removed"
	return nil
}

func failingOnlyNonRetryableError_v2() *resource.RetryError { // want "RetryFunc should include RetryableError\\(\\) handling or be removed"
	return resource.NonRetryableError(errors.New(""))
}

func passingAnonymousRetryFunc_v2() {
	_ = resource.Retry(1*time.Minute, func() *resource.RetryError {
		return resource.RetryableError(errors.New(""))
	})
}

func passingMultipleCallExpr_v2() *resource.RetryError {
	_ = resource.RetryableError(errors.New(""))
	_ = resource.NonRetryableError(errors.New(""))

	return nil
}

func passingRetryableError_v2() *resource.RetryError {
	return resource.RetryableError(errors.New(""))
}
