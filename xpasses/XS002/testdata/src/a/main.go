package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{ // want "schema attributes should be in alphabetical order"
			"name": {},
			"arn":  {},
		},
	}

	_ = schema.Resource{
		Schema: map[string]*schema.Schema{ // want "schema attributes should be in alphabetical order"
			"name": {},
			"arn":  {},
			"block": {},
		},
	}

	_ = schema.Resource{
		Schema: map[string]*schema.Schema{ // want "schema attributes should be in alphabetical order"
			"block": {},
			"arn":  {},
			"name": {},
		},
	}

	_ = schema.Resource{
		Schema: map[string]*schema.Schema{ // want "schema attributes should be in alphabetical order"
			"arn":  {},
			"name": {},
			"block": {},
		},
	}

	// safely ignore const references in schema
	const myConstKey = "my_const_key"
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			myConstKey: {},
			"arn": {},
			"name":  {},
		},
	}

	// safely ignore var references in schema
	var myVarKey = "my_var_key"
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			myVarKey: {},
			"arn": {},
			"name":  {},
		},
	}

	// safely ignore short var declaration references in schema
	myAssignKey := "my_assign_key"
	_ = schema.Resource{
		Schema: map[string]*schema.Schema{
			myAssignKey: {},
			"arn": {},
			"name":  {},
		},
	}
}
