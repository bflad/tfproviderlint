package a

import (
	s "github.com/hashicorp/terraform/helper/schema"
)

func fcommentignore() {
	_ = s.Resource{Exists: existsFunc} //lintignore:resourceexistsfunc

	//lintignore:resourceexistsfunc
	_ = s.Resource{Exists: existsFunc}
}
