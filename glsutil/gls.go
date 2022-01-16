package glsutil

import (
	"context"

	"github.com/jtolds/gls"
)

const glsContextKey = "gls_context_key"

// GlsSetContext 向gls设置ctx
func GlsSetContext(ctx context.Context, cb func()) {
	mgr.SetValues(gls.Values{glsContextKey: ctx}, cb)
}

// GlsContext 从gls获取ctx
func GlsContext() (ctx context.Context, ctxIsDefault bool) {
	glsCtx, ok := mgr.GetValue(glsContextKey)
	if ok {
		ctx = glsCtx.(context.Context)
	} else {
		ctx = context.Background()
		ctxIsDefault = true
	}
	return
}
