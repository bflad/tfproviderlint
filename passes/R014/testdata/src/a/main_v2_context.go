package a

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Comment ignored

//lintignore:R014
func commentIgnoreFirstParameter_v2_context(ctx context.Context, invalid *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

//lintignore:R014
func commentIgnoreSecondParameter_v2_context(ctx context.Context, d *schema.ResourceData, invalid interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

// Failing

func failingAnonymousFirstParameter_v2_context() {
	_ = func(ctx context.Context, invalid *schema.ResourceData, meta interface{}) diag.Diagnostics { // want "\\*schema.ResourceData parameter of CreateContext, ReadContext, UpdateContext, or DeleteContext should be named d"
		return diag.Diagnostics{}
	}
}

func failingAnonymousSecondParameter_v2_context() {
	_ = func(ctx context.Context, d *schema.ResourceData, invalid interface{}) diag.Diagnostics { // want "interface\\{\\} parameter of CreateContext, ReadContext, UpdateContext, or DeleteContext should be named meta"
		return diag.Diagnostics{}
	}
}

func failingFirstParameter_v2_context(ctx context.Context, invalid *schema.ResourceData, meta interface{}) diag.Diagnostics { // want "\\*schema.ResourceData parameter of CreateContext, ReadContext, UpdateContext, or DeleteContext should be named d"
	return diag.Diagnostics{}
}

func failingSecondParameter_v2_context(ctx context.Context, d *schema.ResourceData, invalid interface{}) diag.Diagnostics { // want "interface\\{\\} parameter of CreateContext, ReadContext, UpdateContext, or DeleteContext should be named meta"
	return diag.Diagnostics{}
}

// Passing

func passingAnonymous_v2_context() {
	_ = func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
		return diag.Diagnostics{}
	}
}

func passing_v2_context(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func passingOther_v2_context(ctx context.Context, diff *schema.ResourceDiff, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}
