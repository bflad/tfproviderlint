package a

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	var d schema.ResourceData

	d.SetId(time.Now().Format(time.RFC3339)) //lintignore:R017

	//lintignore:R017
	d.SetId(time.Now().Format(time.RFC3339))
}
