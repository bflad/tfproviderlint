package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func f() {
	var d schema.ResourceData
	key := "test"

	d.Set(key, "") // want "ResourceData.Set\\(\\) key argument should be string literal"

	d.Set("test", "")
	fResourceFunc(&d, nil)
}

func fResourceFunc(d *schema.ResourceData, meta interface{}) error {
	key := ""

	d.Set(key, "") // want "ResourceData.Set\\(\\) key argument should be string literal"

	d.Set("test", "")

	return nil
}
