package schema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Schema struct {
	Type    schema.ValueType
	Default interface{}
}
