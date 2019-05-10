package a

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func f() {
	var t *testing.T

	_ = resource.TestCase{} // want "missing CheckDestroy"

	_ = resource.TestCase{
		CheckDestroy: nil,
	}

	resource.Test(t, resource.TestCase{}) // want "missing CheckDestroy"

	resource.Test(t, resource.TestCase{
		CheckDestroy: nil,
	})
}
