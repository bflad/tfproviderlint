package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	_ = s.Resource{Exists: existsFunc} //lintignore:R003

	//lintignore:R003
	_ = s.Resource{Exists: existsFunc}
}
