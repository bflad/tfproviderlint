package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	// Comment ignored

	//lintignore:R012
	_ = schema.Resource{
		CustomizeDiff: customizeDiffFunc,
		Read:          readFunc,
	}

	// Failing

	_ = schema.Resource{ // want "data source should not configure CustomizeDiff"
		CustomizeDiff: customizeDiffFunc,
		Read:          readFunc,
	}

	// Passing

	_ = schema.Resource{
		Create:        createFunc,
		CustomizeDiff: customizeDiffFunc,
		Read:          readFunc,
	}
}

func createFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func customizeDiffFunc(diff *schema.ResourceDiff, v interface{}) error {
	return nil
}

func readFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
