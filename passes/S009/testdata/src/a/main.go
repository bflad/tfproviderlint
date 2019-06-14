package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema of TypeList or TypeSet should not include top level ValidateFunc"
		Type:         schema.TypeList,
		ValidateFunc: validateFunc,
	}

	_ = schema.Schema{ // want "schema of TypeList or TypeSet should not include top level ValidateFunc"
		Type:         schema.TypeSet,
		ValidateFunc: validateFunc,
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
			ValidateFunc: validateFunc,
		},
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema of TypeList or TypeSet should not include top level ValidateFunc"
			Type:         schema.TypeSet,
			ValidateFunc: validateFunc,
		},
	}
}

func validateFunc(v interface{}, k string) (ws []string, errors []error) {
	return ws, errors
}
