package a

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	var d schema.ResourceData

	d.SetId(time.Now().Format(time.RFC3339)) //lintignore:R017

	//lintignore:R017
	d.SetId(time.Now().Format(time.RFC3339))
}
