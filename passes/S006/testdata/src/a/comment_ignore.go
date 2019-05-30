package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func fcommentignore() {
	_ = schema.Schema{Type: schema.TypeMap} //lintignore:S006

	//lintignore:S006
	_ = schema.Schema{
		Type: schema.TypeMap,
	}
}
