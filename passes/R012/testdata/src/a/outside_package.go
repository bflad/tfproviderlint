package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Resource{
		CustomizeDiff: outsideCustomizeDiffFunc,
		Read:          outsideReadFunc,
	}
}

func outsideCustomizeDiffFunc(diff *schema.ResourceDiff, v interface{}) error {
	return nil
}

func outsideReadFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
