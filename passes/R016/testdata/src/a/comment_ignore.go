package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	var d schema.ResourceData

	d.SetId(resource.PrefixedUniqueId("test")) //lintignore:R016

	//lintignore:R016
	d.SetId(resource.PrefixedUniqueId("test"))
}
