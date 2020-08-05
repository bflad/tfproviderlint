package a

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	// Comment ignored

	//lintignore:R012
	_ = schema.Resource{
		CustomizeDiff: customizeDiffFunc_v2,
		Read:          readFunc_v2,
	}

	// Failing

	_ = schema.Resource{ // want "data source should not configure CustomizeDiff"
		CustomizeDiff: customizeDiffFunc_v2,
		Read:          readFunc_v2,
	}

	// Passing

	_ = schema.Resource{
		Create:        createFunc_v2,
		CustomizeDiff: customizeDiffFunc_v2,
		Read:          readFunc_v2,
	}
}

func createFunc_v2(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func customizeDiffFunc_v2(ctx context.Context, diff *schema.ResourceDiff, v interface{}) error {
	return nil
}

func readFunc_v2(d *schema.ResourceData, meta interface{}) error {
	return nil
}
