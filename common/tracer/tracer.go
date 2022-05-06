package tracer

import (
	"context"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
)

const TracerName = "bcsns"

type Config struct {
	Name              string  `yaml:"name,omitempty"`
	Exporter          string  `yaml:"exporter,omitempty"`
	CollectorEndpoint string  `yaml:"collector_endpoint,omitempty"`
	SampleRatio       float64 `yaml:"sample_ratio,omitempty"`
}

func Init(c Config) (cancel func()) {
	var exporter sdktrace.SpanExporter
	var err error

	switch c.Exporter {
	case "stdout":
		exporter, err = stdout.New(stdout.WithPrettyPrint())
	case "jaeger":
		exporter, err = jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(c.CollectorEndpoint)))
	case "zipkin":
		fallthrough
	default:
		panic("not implemented trace exporter")
	}

	if err != nil {
		panic(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(c.SampleRatio))),
		sdktrace.WithBatcher(exporter),
		// resource是用来标识服务信息的，semconv.SchemaURL表示这些信息的版本，后面的参数就是kv形式，ServiceNameKey是设置服务名，key是"service.name"
		sdktrace.WithResource(resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String(c.Name))),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return func() { tp.Shutdown(context.Background()) }
}

var RpcServerTraceInterceptor = otelgrpc.UnaryServerInterceptor()

var RpcClientTraceInterceptor = otelgrpc.UnaryClientInterceptor()

var HttpServerTraceHandler = otelhttp.NewHandler

func Tracer() trace.Tracer {
	return otel.Tracer(TracerName)
}

func WithTraceContext(ctx context.Context, kvs ...string) context.Context {
	md := metadata.Pairs(
		append(kvs, "timestamp", time.Now().Format(time.StampNano))...,
	)

	return metadata.NewOutgoingContext(ctx, md)
}
