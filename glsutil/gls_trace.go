package glsutil

import (
	"context"

	"github.com/jtolds/gls"
)

var mgr = gls.NewContextManager()

const (
	traceIDKey  = "trace_id"
	pSpanIDKey  = "p_span_id"
	spanIDKey   = "span_id"
	traceCtxKey = "ctx_dlog_order_map_key"
)

// SetOpenTracingGls 向gls设置OpenTracing信息
func SetOpenTracingGls(traceID, pSpanID, spanID string, cb func()) {
	mgr.SetValues(gls.Values{traceIDKey: traceID, pSpanIDKey: pSpanID, spanIDKey: spanID}, cb)
}

// SetOpenTracingWithCtx 向gls设置OpenTracing，ctx信息
func SetOpenTracingWithCtx(traceID, pSpanID, spanID string, ctx context.Context, cb func()) {
	mgr.SetValues(gls.Values{traceIDKey: traceID, pSpanIDKey: pSpanID, spanIDKey: spanID, traceCtxKey: ctx}, cb)
}

// GetOpenTracingFromGls 获取gls中OpenTracing信息
func GetOpenTracingFromGls() (traceID string, pSpanID string, spanID string) {
	trace, ok := mgr.GetValue(traceIDKey)
	if ok {
		traceID = trace.(string)
	}
	pSpan, ok := mgr.GetValue(pSpanIDKey)
	if ok {
		pSpanID = pSpan.(string)
	}
	span, ok := mgr.GetValue(spanIDKey)
	if ok {
		spanID = span.(string)
	}
	return
}

// TraceCtx 从gls中获取ctx
func TraceCtx() (ctx context.Context) {
	traceCtx, ok := mgr.GetValue(traceCtxKey)
	if ok {
		ctx = traceCtx.(context.Context)
	} else {
		ctx = context.Background()
	}
	return
}
