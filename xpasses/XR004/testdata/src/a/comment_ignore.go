package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fcommentignore() {
	var d schema.ResourceData

	d.Set("test", []interface{}{})  //lintignore:XR004
	d.Set("test", []string{})       //lintignore:XR004
	d.Set("test", sliceInterface()) //lintignore:XR004

	d.Set("test", map[string]interface{}{}) //lintignore:XR004
	d.Set("test", map[string]string{})      //lintignore:XR004
	d.Set("test", mapInterface())           //lintignore:XR004

	//lintignore:XR004
	d.Set("test", []interface{}{})
	//lintignore:XR004
	d.Set("test", []string{})
	//lintignore:XR004
	d.Set("test", sliceInterface())

	//lintignore:XR004
	d.Set("test", map[string]interface{}{})
	//lintignore:XR004
	d.Set("test", map[string]string{})
	//lintignore:XR004
	d.Set("test", mapInterface())
}
