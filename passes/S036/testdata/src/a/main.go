package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	// Comment ignored

	//lintignore:S036
	_ = schema.Schema{
		ConflictsWith: []string{
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
		ConflictsWith: []string{
			".invalidreference",       // want "invalid ConflictsWith attribute reference"
			"InvalidReference",        // want "invalid ConflictsWith attribute reference"
			"invalidreference!",       // want "invalid ConflictsWith attribute reference"
			"invalid-reference",       // want "invalid ConflictsWith attribute reference"
			"invalid.reference",       // want "invalid ConflictsWith attribute reference"
			"invalid..reference",      // want "invalid ConflictsWith attribute reference"
			"invalid.123.reference",   // want "invalid ConflictsWith attribute reference"
			"invalid.0.sub.reference", // want "invalid ConflictsWith attribute reference"
			"invalid.sub.0.reference", // want "invalid ConflictsWith attribute reference"
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			ConflictsWith: []string{
				".invalidreference",       // want "invalid ConflictsWith attribute reference"
				"InvalidReference",        // want "invalid ConflictsWith attribute reference"
				"invalidreference!",       // want "invalid ConflictsWith attribute reference"
				"invalid-reference",       // want "invalid ConflictsWith attribute reference"
				"invalid.reference",       // want "invalid ConflictsWith attribute reference"
				"invalid..reference",      // want "invalid ConflictsWith attribute reference"
				"invalid.123.reference",   // want "invalid ConflictsWith attribute reference"
				"invalid.0.sub.reference", // want "invalid ConflictsWith attribute reference"
				"invalid.sub.0.reference", // want "invalid ConflictsWith attribute reference"
			},
		},
	}

	// Passing

	_ = schema.Schema{
		ConflictsWith: []string{
			"validreference",
			"valid_reference",
			"valid.0.reference",
			"valid.0.sub.0.reference",
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			ConflictsWith: []string{
				"validreference",
				"valid_reference",
				"valid.0.reference",
				"valid.0.sub.0.reference",
			},
		},
	}
}
