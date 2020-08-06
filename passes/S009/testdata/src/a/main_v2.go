package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	_ = schema.Schema{ // want "schema of TypeList or TypeSet should not include top level ValidateFunc"
		Type:         schema.TypeList,
		ValidateFunc: validateFunc_v2,
	}

	_ = schema.Schema{ // want "schema of TypeList or TypeSet should not include top level ValidateFunc"
		Type:         schema.TypeSet,
		ValidateFunc: validateFunc_v2,
	}

	_ = schema.Schema{
		Type: schema.TypeList,
	}

	_ = schema.Schema{
		Type: schema.TypeSet,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema of TypeList or TypeSet should not include top level ValidateFunc"
			Type:         schema.TypeList,
			ValidateFunc: validateFunc_v2,
		},
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema of TypeList or TypeSet should not include top level ValidateFunc"
			Type:         schema.TypeSet,
			ValidateFunc: validateFunc_v2,
		},
	}
}

func validateFunc_v2(v interface{}, k string) (ws []string, errors []error) {
	return ws, errors
}
