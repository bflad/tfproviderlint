package a

import (
	"testdata/src/a/schema"
)

func foutside() {
	_ = schema.Schema{
		Computed:     true,
		ComputedWhen: []string{"another_attr"},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed:     true,
			ComputedWhen: []string{"another_attr"},
		},
	}
}
