package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type MyBool bool
type MyString string
type MyInt int

func f_v2() {
	_ = schema.Schema{
		Type:    schema.TypeString,
		Default: true, // want "schema should not declare Default with incompatible Type"
	}

	_ = schema.Schema{
		Type:    schema.TypeBool,
		Default: true,
	}
	_ = schema.Schema{
		Type:    schema.TypeBool,
		Default: bool(MyBool(true)),
	}

	_ = schema.Schema{
		Type:    schema.TypeString,
		Default: "foo",
	}
	_ = schema.Schema{
		Type:    schema.TypeString,
		Default: string(MyString("foo")),
	}

	_ = schema.Schema{
		Type:    schema.TypeInt,
		Default: 123,
	}
	_ = schema.Schema{
		Type:    schema.TypeInt,
		Default: int(MyInt(123)),
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type:    schema.TypeInt,
			Default: true, // want "schema should not declare Default with incompatible Type"
		},
	}
}
