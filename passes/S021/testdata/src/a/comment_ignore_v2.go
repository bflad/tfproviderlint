package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:S021
	_ = schema.Schema{
		Computed:     true,
		ComputedWhen: []string{"another_attr"},
	}

	//lintignore:S021
	_ = map[string]*schema.Schema{
		"name": {
			Computed:     true,
			ComputedWhen: []string{"another_attr"},
		},
	}
}
