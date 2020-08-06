package a

import (
	r "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func fcommentignore_v2() {
	_ = r.TestCase{} //lintignore:AT001

	//lintignore:AT001
	_ = r.TestCase{}
}
