package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	var d schema.ResourceData
	
	_ = d.HasChange("attr1") || d.HasChange("attr2") //lintignore:R005

	//lintignore:R005
	_ = d.HasChange("attr1") || d.HasChange("attr2")

	//lintignore:R005
	_ = d.HasChange("attr1") ||
		d.HasChange("attr2")
}
