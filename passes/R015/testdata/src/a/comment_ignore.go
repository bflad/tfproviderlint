package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	var d schema.ResourceData

	d.SetId(resource.UniqueId()) //lintignore:R015

	//lintignore:R015
	d.SetId(resource.UniqueId())
}
