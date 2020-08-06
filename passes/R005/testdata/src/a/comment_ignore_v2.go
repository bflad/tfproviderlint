package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	var d schema.ResourceData

	_ = d.HasChange("attr1") || d.HasChange("attr2") //lintignore:R005

	//lintignore:R005
	_ = d.HasChange("attr1") || d.HasChange("attr2")

	//lintignore:R005
	_ = d.HasChange("attr1") ||
		d.HasChange("attr2")
}
