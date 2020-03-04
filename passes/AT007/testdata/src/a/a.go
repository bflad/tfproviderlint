package a

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccExampleThing_single(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{})
}

//lintignore:AT007
func TestAccExampleThing_ignored(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{})
	resource.ParallelTest(t, resource.TestCase{})
}

func TestAccExampleThing_multiple(t *testing.T) { // want "acceptance test function should contain only one ParallelTest invocation"
	resource.ParallelTest(t, resource.TestCase{})
	resource.ParallelTest(t, resource.TestCase{})
}
