package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	var d schema.ResourceData

	d.Set("test", []interface{}{})     //lintignore:XR004
	d.Set("test", []string{})          //lintignore:XR004
	d.Set("test", sliceInterface_v2()) //lintignore:XR004

	d.Set("test", map[string]interface{}{}) //lintignore:XR004
	d.Set("test", map[string]string{})      //lintignore:XR004
	d.Set("test", mapInterface_v2())        //lintignore:XR004

	//lintignore:XR004
	d.Set("test", []interface{}{})
	//lintignore:XR004
	d.Set("test", []string{})
	//lintignore:XR004
	d.Set("test", sliceInterface_v2())

	//lintignore:XR004
	d.Set("test", map[string]interface{}{})
	//lintignore:XR004
	d.Set("test", map[string]string{})
	//lintignore:XR004
	d.Set("test", mapInterface_v2())
}
