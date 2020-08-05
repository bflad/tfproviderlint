package a

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestUnitTest_v2(t *testing.T) {}

func TestUnitTest_ResourceUnitTest_v2(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{})
}

func TestExampleThing_basic_v2(t *testing.T) { // want "acceptance test function name should begin with TestAcc"
	resource.Test(t, resource.TestCase{})
}

func TestExampleWidget_basic_v2(t *testing.T) { // want "acceptance test function name should begin with TestAcc"
	resource.ParallelTest(t, resource.TestCase{})
}

//lintignore:AT005
func TestExampleThing_ignored_v2(t *testing.T) {
	resource.Test(t, resource.TestCase{})
}

//lintignore:AT005
func TestExampleWidget_ignored_v2(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{})
}

func TestAccExampleThing_basic_v2(t *testing.T) {
	resource.Test(t, resource.TestCase{})
}

func TestAccExampleWidget_basic_v2(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{})
}
