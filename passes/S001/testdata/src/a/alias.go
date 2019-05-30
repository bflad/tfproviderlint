package a

import (
	s "github.com/hashicorp/terraform/helper/schema"
)

func falias() {
	_ = s.Schema{ // want "schema of TypeMap should include Elem"
		Type: s.TypeMap,
	}
}
