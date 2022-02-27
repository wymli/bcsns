package logx

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

func WithTraceCtx(ctx context.Context) Logger {
	l := defaultLogger.With().Str("trace_id", GetTraceIdFromCtx(ctx)).Str("span_id", GetSpanIdFromCtx(ctx)).Logger()
	return Logger{
		Logger: &l,
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
