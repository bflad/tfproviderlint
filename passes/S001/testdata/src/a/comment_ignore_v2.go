package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
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
