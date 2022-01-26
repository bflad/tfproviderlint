package a

import (
	r "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func fcommentignore() {
	_ = r.TestCase{} //lintignore:AT001

	//lintignore:AT001
	_ = r.TestCase{}

	//lintignore:AT001,AT002
	_ = r.TestCase{}

	_ = r.TestCase{} //lintignore:AT001 // extra comment

	//lintignore:AT001 // extra comment
	_ = r.TestCase{}

	//lintignore:AT001,AT002 // extra comment
	_ = r.TestCase{}

	// extra comment
	//lintignore:AT001
	_ = r.TestCase{}

	// extra comment
	//lintignore:AT001
	// extra comment
	_ = r.TestCase{}

	//lintignore:AT001
	// extra comment
	_ = r.TestCase{}
}
