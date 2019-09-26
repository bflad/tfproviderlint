package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	_ = schema.Schema{Type: schema.TypeMap} //lintignore:S006

	//lintignore:S006
	_ = schema.Schema{
		Type: schema.TypeMap,
	}
}
