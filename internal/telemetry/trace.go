package telemetry

import (
	"go.opentelemetry.io/otel/propagation"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semConv "go.opentelemetry.io/otel/semConv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

func New(cfg Trace) trace.Tracer {
	if !cfg.Enabled {
		return trace.NewNoopTracerProvider().Tracer("amirhnajafiz/stan-gee")
	}

	exporter, err := jaeger.New(
		jaeger.WithAgentEndpoint(jaeger.WithAgentHost(cfg.Agent.Host), jaeger.WithAgentPort(cfg.Agent.Port)),
	)
	if err != nil {
		log.Fatalf("failed to initialize export pipline: %v\n", err)
	}

	res, err := resource.Merge(
		resource.Default(),
		resource.NewSchemaless(
			semConv.ServiceNamespaceKey.String("amirhnajafiz"),
			semConv.ServiceNameKey.String("stan-gee"),
		),
	)
	if err != nil {
		panic(err)
	}

	bsp := sdkTrace.NewBatchSpanProcessor(exporter)
	tp := sdkTrace.NewTracerProvider(
		sdkTrace.WithSampler(sdkTrace.ParentBased(sdkTrace.TraceIDRatioBased(cfg.Ratio))),
		sdkTrace.WithSpanProcessor(bsp),
		sdkTrace.WithResource(res),
	)

	otel.SetTracerProvider(tp)

	var tc propagation.TraceContext

	otel.SetTextMapPropagator(tc)

	tracer := otel.Tracer("amirhnajafiz/stan-gee")

	return tracer
}
