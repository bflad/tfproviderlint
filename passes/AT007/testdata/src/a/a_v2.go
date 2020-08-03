package a

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccExampleThing_single_v2(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{})
}

//lintignore:AT007
func TestAccExampleThing_ignored_v2(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{})
	resource.ParallelTest(t, resource.TestCase{})
}

func TestAccExampleThing_multiple_v2(t *testing.T) { // want "acceptance test function should contain only one ParallelTest invocation"
	resource.ParallelTest(t, resource.TestCase{})
	resource.ParallelTest(t, resource.TestCase{})
}
