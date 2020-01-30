package a

import (
	"time"

	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	var d s.ResourceData
	var t time.Time

	d.Set("test", t)                               // want "ResourceData.Set\\(\\) should perform error checking with complex values"
	d.Set("test", []time.Time{t})                  // want "ResourceData.Set\\(\\) should perform error checking with complex values"
	d.Set("test", map[string]time.Time{"test": t}) // want "ResourceData.Set\\(\\) should perform error checking with complex values"

	d.Set("test", nil)

	var b bool

	d.Set("test", b)
	d.Set("test", &b)

	var f32 float32

	d.Set("test", f32)
	d.Set("test", &f32)

	var f64 float64

	d.Set("test", f64)
	d.Set("test", &f64)

	var i int

	d.Set("test", i)
	d.Set("test", &i)

	var i8 int8

	d.Set("test", i8)
	d.Set("test", &i8)

	var i16 int16

	d.Set("test", i16)
	d.Set("test", &i16)

	var i32 int32

	d.Set("test", i32)
	d.Set("test", &i32)

	var i64 int64

	d.Set("test", i64)
	d.Set("test", &i64)

	var s string

	d.Set("test", s)
	d.Set("test", &s)

	d.Set("test", []interface{}{})  // want "ResourceData.Set\\(\\) should perform error checking with complex values"
	d.Set("test", []string{})       // want "ResourceData.Set\\(\\) should perform error checking with complex values"
	d.Set("test", sliceInterface()) // want "ResourceData.Set\\(\\) should perform error checking with complex values"

	d.Set("test", map[string]interface{}{}) // want "ResourceData.Set\\(\\) should perform error checking with complex values"
	d.Set("test", map[string]string{})      // want "ResourceData.Set\\(\\) should perform error checking with complex values"
	d.Set("test", mapInterface())           // want "ResourceData.Set\\(\\) should perform error checking with complex values"

	if err := d.Set("test", []interface{}{}); err != nil {
		return
	}

	if err := d.Set("test", []string{}); err != nil {
		return
	}

	if err := d.Set("test", sliceInterface()); err != nil {
		return
	}

	if err := d.Set("test", map[string]interface{}{}); err != nil {
		return
	}

	if err := d.Set("test", map[string]string{}); err != nil {
		return
	}

	if err := d.Set("test", mapInterface()); err != nil {
		return
	}
}
