package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Schema{
		Elem: &s.Schema{}, // want "schema should not include Elem with incompatible Type"
		Type: s.TypeBool,
	}

	_ = s.Schema{
		Type: s.TypeBool,
	}

	_ = s.Schema{
		Elem: &s.Schema{}, // want "schema should not include Elem with incompatible Type"
		Type: s.TypeFloat,
	}

	_ = s.Schema{
		Type: s.TypeFloat,
	}

	_ = s.Schema{
		Elem: &s.Schema{}, // want "schema should not include Elem with incompatible Type"
		Type: s.TypeInt,
	}

	_ = s.Schema{
		Type: s.TypeInt,
	}

	_ = s.Schema{
		Elem: &s.Schema{}, // want "schema should not include Elem with incompatible Type"
		Type: s.TypeString,
	}

	_ = s.Schema{
		Type: s.TypeString,
	}

	_ = map[string]*s.Schema{
		"name": {
			Elem: &s.Schema{}, // want "schema should not include Elem with incompatible Type"
			Type: s.TypeBool,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Type: s.TypeBool,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Elem: &s.Schema{}, // want "schema should not include Elem with incompatible Type"
			Type: s.TypeFloat,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Type: s.TypeFloat,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Elem: &s.Schema{}, // want "schema should not include Elem with incompatible Type"
			Type: s.TypeInt,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Type: s.TypeInt,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Elem: &s.Schema{}, // want "schema should not include Elem with incompatible Type"
			Type: s.TypeString,
		},
	}

	_ = map[string]*s.Schema{
		"name": {
			Type: s.TypeString,
		},
	}
}
