package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	var d schema.ResourceData

	d.SetId(resource.UniqueId()) //lintignore:R015

	//lintignore:R015
	d.SetId(resource.UniqueId())
}
