package a

import (
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

//lintignore:R006
func commentIgnore() *resource.RetryError {
	return nil
}

func failingAnonymousRetryFunc() {
	_ = resource.Retry(1*time.Minute, func() *resource.RetryError { // want "RetryFunc should include RetryableError\\(\\) handling or be removed"
		return nil
	})
}

func failingNoCallExpr() *resource.RetryError { // want "RetryFunc should include RetryableError\\(\\) handling or be removed"
	return nil
}

func failingOnlyNonRetryableError() *resource.RetryError { // want "RetryFunc should include RetryableError\\(\\) handling or be removed"
	return resource.NonRetryableError(errors.New(""))
}

func passingAnonymousRetryFunc() {
	_ = resource.Retry(1*time.Minute, func() *resource.RetryError {
		return resource.RetryableError(errors.New(""))
	})
}

func passingMultipleCallExpr() *resource.RetryError {
	_ = resource.RetryableError(errors.New(""))
	_ = resource.NonRetryableError(errors.New(""))

	return nil
}

func passingRetryableError() *resource.RetryError {
	return resource.RetryableError(errors.New(""))
}
