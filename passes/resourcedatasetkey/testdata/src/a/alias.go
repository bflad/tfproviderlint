package a

import (
	s "github.com/hashicorp/terraform/helper/schema"
)

func falias() {
	var d s.ResourceData
	key := "test"

	d.Set(key, "") // want "ResourceData.Set\\(\\) key argument should be string literal"

	d.Set("test", "")
}
