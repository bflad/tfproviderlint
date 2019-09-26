package a

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
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
