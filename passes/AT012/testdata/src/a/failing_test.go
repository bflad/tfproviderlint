// want "file contains multiple acceptance test name prefixes"
package a

import (
	"testing"
)

func TestAccExampleThingFailing1_Test(t *testing.T) {}

func TestAccExampleThingFailing2_Test(t *testing.T) {}
