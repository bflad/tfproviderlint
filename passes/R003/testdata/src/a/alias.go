package a

import (
	s "github.com/hashicorp/terraform/helper/schema"
)

func falias() {
	_ = s.Resource{Exists: existsFunc} // want "resource should not include Exists function"
}
