package a

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	var d schema.ResourceData
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

	d.Set("test", []interface{}{})     // want "ResourceData.Set\\(\\) should perform error checking with complex values"
	d.Set("test", []string{})          // want "ResourceData.Set\\(\\) should perform error checking with complex values"
	d.Set("test", sliceInterface_v2()) // want "ResourceData.Set\\(\\) should perform error checking with complex values"

	d.Set("test", map[string]interface{}{}) // want "ResourceData.Set\\(\\) should perform error checking with complex values"
	d.Set("test", map[string]string{})      // want "ResourceData.Set\\(\\) should perform error checking with complex values"
	d.Set("test", mapInterface_v2())        // want "ResourceData.Set\\(\\) should perform error checking with complex values"

	if err := d.Set("test", []interface{}{}); err != nil {
		return
	}

	if err := d.Set("test", []string{}); err != nil {
		return
	}

	if err := d.Set("test", sliceInterface_v2()); err != nil {
		return
	}

	if err := d.Set("test", map[string]interface{}{}); err != nil {
		return
	}

	if err := d.Set("test", map[string]string{}); err != nil {
		return
	}

	if err := d.Set("test", mapInterface_v2()); err != nil {
		return
	}

	fResourceFunc_v2(&d, nil)
}

func fResourceFunc_v2(d *schema.ResourceData, meta interface{}) error {
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

	d.Set("test", []interface{}{})     // want "ResourceData.Set\\(\\) should perform error checking with complex values"
	d.Set("test", []string{})          // want "ResourceData.Set\\(\\) should perform error checking with complex values"
	d.Set("test", sliceInterface_v2()) // want "ResourceData.Set\\(\\) should perform error checking with complex values"

	d.Set("test", map[string]interface{}{}) // want "ResourceData.Set\\(\\) should perform error checking with complex values"
	d.Set("test", map[string]string{})      // want "ResourceData.Set\\(\\) should perform error checking with complex values"
	d.Set("test", mapInterface_v2())        // want "ResourceData.Set\\(\\) should perform error checking with complex values"

	if err := d.Set("test", []interface{}{}); err != nil {
		return err
	}

	if err := d.Set("test", []string{}); err != nil {
		return err
	}

	if err := d.Set("test", sliceInterface_v2()); err != nil {
		return err
	}

	if err := d.Set("test", map[string]interface{}{}); err != nil {
		return err
	}

	if err := d.Set("test", map[string]string{}); err != nil {
		return err
	}

	if err := d.Set("test", mapInterface_v2()); err != nil {
		return err
	}

	return nil
}

func mapInterface_v2() map[string]interface{} {
	var mI map[string]interface{}
	return mI
}

func sliceInterface_v2() []interface{} {
	var sI []interface{}
	return sI
}
