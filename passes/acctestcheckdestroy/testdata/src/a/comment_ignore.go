package a

import (
	r "github.com/hashicorp/terraform/helper/resource"
)

func fcommentignore() {
	_ = r.TestCase{} //lintignore:acctestcheckdestroy

	//lintignore:acctestcheckdestroy
	_ = r.TestCase{}
}
