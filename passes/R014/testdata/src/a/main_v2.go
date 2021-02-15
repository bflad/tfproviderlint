package a

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Comment ignored

//lintignore:R014
func commentIgnoreFirstParameter_v2(invalid *schema.ResourceData, meta interface{}) error {
	return nil
}

//lintignore:R014
func commentIgnoreFirstParameter_v2_context(ctx context.Context, invalid *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

//lintignore:R014
func commentIgnoreSecondParameter_v2(d *schema.ResourceData, invalid interface{}) error {
	return nil
}

//lintignore:R014
func commentIgnoreSecondParameter_v2_context(ctx context.Context, d *schema.ResourceData, invalid interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

// Failing

func failingAnonymousFirstParameter_v2() {
	_ = func(invalid *schema.ResourceData, meta interface{}) error { // want "\\*schema.ResourceData parameter of CreateFunc, ReadFunc, UpdateFunc, or DeleteFunc should be named d"
		return nil
	}
}

func failingAnonymousFirstParameter_v2_context() {
	_ = func(ctx context.Context, invalid *schema.ResourceData, meta interface{}) diag.Diagnostics { // want "\\*schema.ResourceData parameter of CreateContextFunc, ReadContextFunc, UpdateContextFunc, or DeleteContextFunc should be named d"
		return diag.Diagnostics{}
	}
}

func failingAnonymousSecondParameter_v2() {
	_ = func(d *schema.ResourceData, invalid interface{}) error { // want "interface\\{\\} parameter of CreateFunc, ReadFunc, UpdateFunc, or DeleteFunc should be named meta"
		return nil
	}
}

func failingAnonymousSecondParameter_v2_context() {
	_ = func(ctx context.Context, d *schema.ResourceData, invalid interface{}) diag.Diagnostics { // want "interface\\{\\} parameter of CreateContextFunc, ReadContextFunc, UpdateContextFunc, or DeleteContextFunc should be named meta"
		return diag.Diagnostics{}
	}
}

func failingFirstParameter_v2(invalid *schema.ResourceData, meta interface{}) error { // want "\\*schema.ResourceData parameter of CreateFunc, ReadFunc, UpdateFunc, or DeleteFunc should be named d"
	return nil
}

func failingFirstParameter_v2_context(ctx context.Context, invalid *schema.ResourceData, meta interface{}) diag.Diagnostics { // want "\\*schema.ResourceData parameter of CreateContextFunc, ReadContextFunc, UpdateContextFunc, or DeleteContextFunc should be named d"
	return diag.Diagnostics{}
}

func failingSecondParameter_v2(d *schema.ResourceData, invalid interface{}) error { // want "interface\\{\\} parameter of CreateFunc, ReadFunc, UpdateFunc, or DeleteFunc should be named meta"
	return nil
}

func failingSecondParameter_v2_context(ctx context.Context, d *schema.ResourceData, invalid interface{}) diag.Diagnostics { // want "interface\\{\\} parameter of CreateContextFunc, ReadContextFunc, UpdateContextFunc, or DeleteContextFunc should be named meta"
	return diag.Diagnostics{}
}

// Passing

func passingAnonymous_v2() {
	_ = func(d *schema.ResourceData, meta interface{}) error {
		return nil
	}
}

func passingAnonymous_v2_context() {
	_ = func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
		return diag.Diagnostics{}
	}
}

func passing_v2(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func passing_v2_context(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func passingOther_v2(diff *schema.ResourceDiff, meta interface{}) error {
	return nil
}

func passingOther_v2_context(ctx context.Context, diff *schema.ResourceDiff, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func passingExtended_v2(invalidResourceData *schema.ResourceData, invalidInterface interface{}, notCRUDFunc string) error {
	return nil
}

func passingExtended_v2_context(invalidContext context.Context, invalidResourceData *schema.ResourceData, invalidInterface interface{}, notCRUDContextFunc string) diag.Diagnostics {
	return diag.Diagnostics{}
}
