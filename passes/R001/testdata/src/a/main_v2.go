package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	var d schema.ResourceData
	key := "test"

	d.Set(key, "") // want "ResourceData.Set\\(\\) key argument should be string literal"

	d.Set("test", "")
	fResourceFunc_v2(&d, nil)
}

func fResourceFunc_v2(d *schema.ResourceData, meta interface{}) error {
	key := ""

	d.Set(key, "") // want "ResourceData.Set\\(\\) key argument should be string literal"

	d.Set("test", "")

	return nil
}
