package a

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f_v2() {
	var d schema.ResourceData
	var t time.Time

	d.Set("test", t)                               // want "ResourceData.Set\\(\\) incompatible value type"
	d.Set("test", []time.Time{t})                  // want "ResourceData.Set\\(\\) incompatible value type"
	d.Set("test", map[string]time.Time{"test": t}) // want "ResourceData.Set\\(\\) incompatible value type"

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

	d.Set("test", []interface{}{})
	d.Set("test", []string{})

	d.Set("test", map[string]interface{}{})
	d.Set("test", map[string]string{})

	fResourceFunc_v2(&d, nil)
}

func fResourceFunc_v2(d *schema.ResourceData, meta interface{}) error {
	var t time.Time

	d.Set("test", t)                               // want "ResourceData.Set\\(\\) incompatible value type"
	d.Set("test", []time.Time{t})                  // want "ResourceData.Set\\(\\) incompatible value type"
	d.Set("test", map[string]time.Time{"test": t}) // want "ResourceData.Set\\(\\) incompatible value type"

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

	d.Set("test", []interface{}{})
	d.Set("test", []string{})
	d.Set("test", sliceInterface())

	d.Set("test", map[string]interface{}{})
	d.Set("test", map[string]string{})
	d.Set("test", mapInterface())
	d.Set("test", mapInterface()["doesnotexist"])

	return nil
}

func mapInterface() map[string]interface{} {
	var mI map[string]interface{}
	return mI
}

func sliceInterface() []interface{} {
	var sI []interface{}
	return sI
}
