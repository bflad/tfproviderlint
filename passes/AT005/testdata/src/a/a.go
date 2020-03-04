package a

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestUnitTest(t *testing.T) {}

func TestUnitTest_ResourceUnitTest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{})
}

func TestExampleThing_basic(t *testing.T) { // want "acceptance test function name should begin with TestAcc"
	resource.Test(t, resource.TestCase{})
}

func TestExampleWidget_basic(t *testing.T) { // want "acceptance test function name should begin with TestAcc"
	resource.ParallelTest(t, resource.TestCase{})
}

//lintignore:AT005
func TestExampleThing_ignored(t *testing.T) {
	resource.Test(t, resource.TestCase{})
}

//lintignore:AT005
func TestExampleWidget_ignored(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{})
}

func TestAccExampleThing_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{})
}

func TestAccExampleWidget_basic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{})
}
