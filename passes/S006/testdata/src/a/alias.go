package a

import (
	"github.com/hashicorp/terraform/helper/schema"
	s "github.com/hashicorp/terraform/helper/schema"
)

func falias() {
	_ = s.Schema{ // want "schema of TypeMap should include Elem"
		Type: s.TypeMap,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "schema of TypeMap should include Elem"
			Type: s.TypeMap,
		},
	}
}
