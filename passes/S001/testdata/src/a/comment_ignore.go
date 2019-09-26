package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	_ = schema.Schema{Type: schema.TypeList} //lintignore:S001

	_ = schema.Schema{Type: schema.TypeSet} //lintignore:S001

	//lintignore:S001
	_ = schema.Schema{
		Type: schema.TypeList,
	}

	//lintignore:S001
	_ = schema.Schema{
		Type: schema.TypeSet,
	}
}
