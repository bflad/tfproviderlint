package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func falias_v2() {
	var d s.ResourceData
	key := "test"

	d.Set(key, "") // want "ResourceData.Set\\(\\) key argument should be string literal"

	d.Set("test", "")
}
