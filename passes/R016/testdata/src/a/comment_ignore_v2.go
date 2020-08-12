package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	var d schema.ResourceData

	d.SetId(resource.PrefixedUniqueId("test")) //lintignore:R016

	//lintignore:R016
	d.SetId(resource.PrefixedUniqueId("test"))
}
