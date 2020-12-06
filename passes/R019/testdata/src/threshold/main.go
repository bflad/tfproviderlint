package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	var d schema.ResourceData

	// Failing with lower threshold

	d.HasChanges("attr1", "attr2", "attr3") // want "d.HasChanges\\(\\) has many arguments, consider d.HasChangesExcept\\(\\)"

	// Passing with lower threshold

	d.HasChanges("attr1", "attr2")
}
