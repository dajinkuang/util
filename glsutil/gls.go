package glsutil

import (
	"context"

	"github.com/jtolds/gls"
)

var (
	glsContextKey = "gls_context_key"
)

func GlsSetContext(ctx context.Context, cb func()) {
	mgr.SetValues(gls.Values{glsContextKey: ctx}, cb)
}

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
