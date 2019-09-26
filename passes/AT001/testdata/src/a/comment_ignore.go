package a

import (
	r "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func fcommentignore() {
	_ = r.TestCase{} //lintignore:AT001

	//lintignore:AT001
	_ = r.TestCase{}
}
