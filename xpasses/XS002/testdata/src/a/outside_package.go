package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {},
			"arn": {},
		},
	}
}



