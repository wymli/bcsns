package logx

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type KvEntry struct {
	Key   string
	Value interface{}
}

func WithTraceCtx(ctx context.Context, kvs ...KvEntry) Logger {
	l := GlobalLogger.With().Str("trace_id", GetTraceIdFromCtx(ctx)).Str("span_id", GetSpanIdFromCtx(ctx))
	for _, kv := range kvs {
		l = l.Interface(kv.Key, kv.Value)
	}
	ll := l.Logger()

	return Logger{
		Logger: &ll,
	}
}

func GetSpanIdFromCtx(ctx context.Context) string {
	spanCtx := trace.SpanContextFromContext(ctx)
	if !spanCtx.HasSpanID() {
		return ""
	}

	return spanCtx.SpanID().String()
}

func GetTraceIdFromCtx(ctx context.Context) string {
	spanCtx := trace.SpanContextFromContext(ctx)
	if !spanCtx.HasTraceID() {
		return ""
	}

	return spanCtx.TraceID().String()
}

func GetSpanCtxFromCtx(ctx context.Context) trace.SpanContext {
	return trace.SpanContextFromContext(ctx)
}
