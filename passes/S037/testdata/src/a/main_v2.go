package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	// Comment ignored

	//lintignore:S037
	_ = schema.Schema{
		ExactlyOneOf: []string{
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
		ExactlyOneOf: []string{
			".invalidreference",       // want "invalid ExactlyOneOf attribute reference"
			"InvalidReference",        // want "invalid ExactlyOneOf attribute reference"
			"invalidreference!",       // want "invalid ExactlyOneOf attribute reference"
			"invalid-reference",       // want "invalid ExactlyOneOf attribute reference"
			"invalid.reference",       // want "invalid ExactlyOneOf attribute reference"
			"invalid..reference",      // want "invalid ExactlyOneOf attribute reference"
			"invalid.123.reference",   // want "invalid ExactlyOneOf attribute reference"
			"invalid.0.sub.reference", // want "invalid ExactlyOneOf attribute reference"
			"invalid.sub.0.reference", // want "invalid ExactlyOneOf attribute reference"
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			ExactlyOneOf: []string{
				".invalidreference",       // want "invalid ExactlyOneOf attribute reference"
				"InvalidReference",        // want "invalid ExactlyOneOf attribute reference"
				"invalidreference!",       // want "invalid ExactlyOneOf attribute reference"
				"invalid-reference",       // want "invalid ExactlyOneOf attribute reference"
				"invalid.reference",       // want "invalid ExactlyOneOf attribute reference"
				"invalid..reference",      // want "invalid ExactlyOneOf attribute reference"
				"invalid.123.reference",   // want "invalid ExactlyOneOf attribute reference"
				"invalid.0.sub.reference", // want "invalid ExactlyOneOf attribute reference"
				"invalid.sub.0.reference", // want "invalid ExactlyOneOf attribute reference"
			},
		},
	}

	// Passing

	_ = schema.Schema{
		ExactlyOneOf: []string{
			"validreference",
			"valid_reference",
			"valid.0.reference",
			"valid.0.sub.0.reference",
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			ExactlyOneOf: []string{
				"validreference",
				"valid_reference",
				"valid.0.reference",
				"valid.0.sub.0.reference",
			},
		},
	}
}
