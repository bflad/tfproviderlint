package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	_ = s.Resource{Exists: existsFunc_v2} //lintignore:R003

	//lintignore:R003
	_ = s.Resource{Exists: existsFunc_v2}
}
