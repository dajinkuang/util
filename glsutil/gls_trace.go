package glsutil

import (
	"context"
	"github.com/jtolds/gls"
)

var (
	mgr         = gls.NewContextManager()
	traceIDKey  = "trace_id"
	pSpanIDKey  = "p_span_id"
	spanIDKey   = "span_id"
	traceCtxKey = "ctx_dlog_order_map_key"
)

func SetOpenTracingGls(traceID, pSpanID, spanID string, cb func()) {
	mgr.SetValues(gls.Values{traceIDKey: traceID, pSpanIDKey: pSpanID, spanIDKey: spanID}, cb)
}

func SetOpenTracingWithCtx(traceID, pSpanID, spanID string, ctx context.Context, cb func()) {
	mgr.SetValues(gls.Values{traceIDKey: traceID, pSpanIDKey: pSpanID, spanIDKey: spanID, traceCtxKey: ctx}, cb)
}

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

func TraceCtx() (ctx context.Context) {
	traceCtx, ok := mgr.GetValue(traceCtxKey)
	if ok {
		ctx = traceCtx.(context.Context)
	} else {
		ctx = context.Background()
	}
	return
}
