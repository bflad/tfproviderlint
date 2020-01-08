package a

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestUnitTest(t *testing.T) {}

func TestExampleThing_basic(t *testing.T) { // want "acceptance test function name should begin with TestAcc"
	resource.Test(t, resource.TestCase{})
}

func TestExampleWidget_basic(t *testing.T) { // want "acceptance test function name should begin with TestAcc"
	resource.ParallelTest(t, resource.TestCase{})
}

func TestAccExampleThing_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{})
}

func TestAccExampleWidget_basic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{})
}
