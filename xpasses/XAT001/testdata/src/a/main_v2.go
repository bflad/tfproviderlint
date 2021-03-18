package a

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func f_v2() {
	var t *testing.T

	// Failing

	_ = resource.TestCase{} // want "missing ErrorCheck"

	resource.Test(t, resource.TestCase{}) // want "missing ErrorCheck"

	// Comment Ignored

	_ = resource.TestCase{} //lintignore:XAT001

	//lintignore:XAT001
	_ = resource.TestCase{}

	// Passing

	_ = resource.TestCase{
		ErrorCheck: nil,
	}

	resource.Test(t, resource.TestCase{
		ErrorCheck: nil,
	})
}
