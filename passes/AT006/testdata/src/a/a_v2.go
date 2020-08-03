package a

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccExampleThing_single_v2(t *testing.T) {
	resource.Test(t, resource.TestCase{})
}

//lintignore:AT006
func TestAccExampleThing_ignored_v2(t *testing.T) {
	resource.Test(t, resource.TestCase{})
	resource.Test(t, resource.TestCase{})
}

func TestAccExampleThing_multiple_v2(t *testing.T) { // want "acceptance test function should contain only one Test invocation"
	resource.Test(t, resource.TestCase{})
	resource.Test(t, resource.TestCase{})
}
