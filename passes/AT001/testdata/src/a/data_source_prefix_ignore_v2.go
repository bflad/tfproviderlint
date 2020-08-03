package a

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func ffilenameignore_v2() {
	var t *testing.T

	_ = resource.TestCase{}

	resource.Test(t, resource.TestCase{})
}
