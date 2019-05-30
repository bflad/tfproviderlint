package a

import (
	r "github.com/hashicorp/terraform/helper/resource"
)

func fcommentignore() {
	_ = r.TestCase{} //lintignore:AT001

	//lintignore:AT001
	_ = r.TestCase{}
}
