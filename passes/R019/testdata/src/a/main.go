package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	var d schema.ResourceData

	// Comment ignored

	//lintignore:R019
	d.HasChanges("attr1", "attr2", "attr3", "attr4", "attr5")

	d.HasChanges("attr1", "attr2", "attr3", "attr4", "attr5") //lintignore:R019

	// Failing

	d.HasChanges("attr1", "attr2", "attr3", "attr4", "attr5") // want "d.HasChanges\\(\\) has many arguments, consider d.HasChangesExcept\\(\\)"

	// Passing

	d.HasChanges("attr1", "attr2", "attr3", "attr4")
}
