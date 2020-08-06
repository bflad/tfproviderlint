package a

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func f_v2() {
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
