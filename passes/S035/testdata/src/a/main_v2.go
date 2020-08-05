package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	// Comment ignored

	//lintignore:S035
	_ = schema.Schema{
		AtLeastOneOf: []string{
			".invalidreference",
			"InvalidReference",
			"invalidreference!",
			"invalid-reference",
			"invalid.reference",
			"invalid..reference",
			"invalid.123.reference",
			"invalid.0.sub.reference",
			"invalid.sub.0.reference",
		},
	}

	// Failing

	_ = schema.Schema{
		AtLeastOneOf: []string{
			".invalidreference",       // want "invalid AtLeastOneOf attribute reference"
			"InvalidReference",        // want "invalid AtLeastOneOf attribute reference"
			"invalidreference!",       // want "invalid AtLeastOneOf attribute reference"
			"invalid-reference",       // want "invalid AtLeastOneOf attribute reference"
			"invalid.reference",       // want "invalid AtLeastOneOf attribute reference"
			"invalid..reference",      // want "invalid AtLeastOneOf attribute reference"
			"invalid.123.reference",   // want "invalid AtLeastOneOf attribute reference"
			"invalid.0.sub.reference", // want "invalid AtLeastOneOf attribute reference"
			"invalid.sub.0.reference", // want "invalid AtLeastOneOf attribute reference"
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			AtLeastOneOf: []string{
				".invalidreference",       // want "invalid AtLeastOneOf attribute reference"
				"InvalidReference",        // want "invalid AtLeastOneOf attribute reference"
				"invalidreference!",       // want "invalid AtLeastOneOf attribute reference"
				"invalid-reference",       // want "invalid AtLeastOneOf attribute reference"
				"invalid.reference",       // want "invalid AtLeastOneOf attribute reference"
				"invalid..reference",      // want "invalid AtLeastOneOf attribute reference"
				"invalid.123.reference",   // want "invalid AtLeastOneOf attribute reference"
				"invalid.0.sub.reference", // want "invalid AtLeastOneOf attribute reference"
				"invalid.sub.0.reference", // want "invalid AtLeastOneOf attribute reference"
			},
		},
	}

	// Passing

	_ = schema.Schema{
		AtLeastOneOf: []string{
			"validreference",
			"valid_reference",
			"valid.0.reference",
			"valid.0.sub.0.reference",
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			AtLeastOneOf: []string{
				"validreference",
				"valid_reference",
				"valid.0.reference",
				"valid.0.sub.0.reference",
			},
		},
	}
}
