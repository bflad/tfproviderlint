package a

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func ffilenameignore() {
	var t *testing.T

	_ = resource.TestCase{}

	resource.Test(t, resource.TestCase{})
}
