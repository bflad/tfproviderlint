package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	_ = schema.Schema{Type: schema.TypeMap} //lintignore:S006

	//lintignore:S006
	_ = schema.Schema{
		Type: schema.TypeMap,
	}
}
