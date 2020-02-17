package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	// Comment ignored

	//lintignore:S034
	_ = schema.Schema{
		PromoteSingle: true,
	}

	// Failing

	_ = schema.Schema{ // want "schema should not enable PromoteSingle"
		PromoteSingle: true,
	}

	_ = schema.Schema{ // want "schema should not enable PromoteSingle"
		PromoteSingle: false,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not enable PromoteSingle"
			PromoteSingle: true,
		},
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema should not enable PromoteSingle"
			PromoteSingle: false,
		},
	}

	// Passing

	_ = schema.Schema{}
}
