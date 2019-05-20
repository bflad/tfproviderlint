package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func fcommentignore() {
	_ = schema.Schema{Type: schema.TypeMap} //lintignore:schematypemapelem

	//lintignore:schematypemapelem
	_ = schema.Schema{
		Type: schema.TypeMap,
	}
}
