package tracing

import (
	"context"

	"github.com/hoojos/target/middleware"
	"github.com/valyala/fasthttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type options struct {
	tracer     trace.Tracer
	propagator propagation.TextMapPropagator
}

type Option func(*options)

func WithTracer(tracer trace.Tracer) Option {
	return func(o *options) {
		o.tracer = tracer
	}
}

func WithPropagator(propagator propagation.TextMapPropagator) Option {
	return func(o *options) {
		o.propagator = propagator
	}
}

func Tracing(opts ...Option) middleware.Middleware {
	var options options
	for _, opt := range opts {
		opt(&options)
	}

	if options.propagator != nil {
		otel.SetTextMapPropagator(options.propagator)
	} else {
		propagator := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
		otel.SetTextMapPropagator(propagator)
	}

	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			_, span := options.tracer.Start(context.Background(), "target")
			defer span.End()
			next(ctx)
		}
	}
}
