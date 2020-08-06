package a

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	var d schema.ResourceData
	var t time.Time

	d.Set("test", t)                               //lintignore:R004
	d.Set("test", []time.Time{t})                  //lintignore:R004
	d.Set("test", map[string]time.Time{"test": t}) //lintignore:R004

	//lintignore:R004
	d.Set("test", t)
	//lintignore:R004
	d.Set("test", []time.Time{t})
	//lintignore:R004
	d.Set("test", map[string]time.Time{"test": t})
}
