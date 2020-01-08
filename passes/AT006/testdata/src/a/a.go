package a

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccExampleThing_single(t *testing.T) {
	resource.Test(t, resource.TestCase{})
}

func TestAccExampleThing_multiple(t *testing.T) { // want "acceptance test function should contain only one Test invocation"
	resource.Test(t, resource.TestCase{})
	resource.Test(t, resource.TestCase{})
}
