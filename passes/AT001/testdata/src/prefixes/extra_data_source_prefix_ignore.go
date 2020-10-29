package prefixes

import (
"testing"

"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func fextraprefixignore() {
	var t *testing.T

	_ = resource.TestCase{}

	resource.Test(t, resource.TestCase{})
}
