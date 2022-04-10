package xtrace

import (
	"context"
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

// TracerType denotes the underlining implementation of opentracing-compatible tracer
type TracerType string

// Tracer is a simple, thin interface for Span creation and SpanContext
// propagation
type Tracer = opentracing.Tracer

// SpanContext represents propagated span identity and state
type SpanContext = jaeger.SpanContext

// StartSpanOption instances (zero or more) may be passed to Tracer.StartSpan.
//
// StartSpanOption borrows from the "functional options" pattern, per
// http://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
type StartSpanOption = opentracing.StartSpanOption

// Span represents an active, un-finished span in the OpenTracing system.
//
// Spans are created by the Tracer interface.
type Span = opentracing.Span

// TracerTypeJaeger identity the Jaeger's tracer implementation
const TracerTypeJaeger TracerType = "jaeger"

// InitDefaultTracer will initialize the default tracer, which is now the Jaeger tracer.
func InitDefaultTracer(serviceName string) error {
	return InitTracer(TracerTypeJaeger, serviceName)
}

// InitTracer provides a way of initialize a customized tracer, which support only the Jaeger tracer currently
func InitTracer(tracerType TracerType, serviceName string) error {
	if bt != nil {
		return nil
	}

	switch tracerType {
	case TracerTypeJaeger:
		return initJaeger(serviceName)
	default:
		return fmt.Errorf("unknown tracer type %v", tracerType)
	}
}

// CloseTracer stop a tracer from collecting trace information, usually this function
//   should be invoked in an graceful exit/handling.
func CloseTracer() error {
	if bt != nil && bt.closer != nil {
		return bt.closer.Close()
	}
	return nil
}

func initJaeger(serviceName string) error {
	configManager := newTracerConfigManager()
	tracerConfig := configManager.GetConfig(serviceName, TracerTypeJaeger)

	cfg, ok := tracerConfig.Payload.(config.Configuration)
	if !ok {
		return fmt.Errorf("imcompatible tracer config %v for jaeger", cfg)
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return err
	}

	bt = &backTracer{
		tracer: tracer,
		closer: closer,
	}

	opentracing.SetGlobalTracer(tracer)
	// TODO: use xlog
	fmt.Printf("init jaeger for %s [done]", serviceName)
	return nil
}

// TODO: perhaps we should make it thread-safe later, if necessary
// singleton
var bt *backTracer

type backTracer struct {
	tracer opentracing.Tracer
	closer io.Closer
}

// String adds a string-valued key:value pair to a Span.LogFields() record
func String(key, value string) log.Field {
	return log.String(key, value)
}

// Int adds an int-valued key:value pair to a Span.LogFields() record
func Int(key string, value int) log.Field {
	return log.Int(key, value)
}

// SpanFromContext returns the `Span` previously associated with `ctx`, or
// `nil` if no such `Span` could be found.
//
// NOTE: context.Context != SpanContext: the former is Go's intra-process
// context propagation mechanism, and the latter houses OpenTracing's per-Span
// identity and baggage information.
func SpanFromContext(ctx context.Context) opentracing.Span {
	return opentracing.SpanFromContext(ctx)
}

// StartSpanFromContext starts and returns a Span with `operationName`, using
// any Span found within `ctx` as a ChildOfRef. If no such parent could be
// found, StartSpanFromContext creates a root (parentless) Span.
//
// The second return value is a context.Context object built around the
// returned Span.
//
// Example usage:
//
//    SomeFunction(ctx context.Context, ...) {
//        sp, ctx := opentracing.StartSpanFromContext(ctx, "SomeFunction")
//        defer sp.Finish()
//        ...
//    }
func StartSpanFromContext(ctx context.Context, operationName string, opts ...StartSpanOption) (Span, context.Context) {
	return opentracing.StartSpanFromContext(ctx, operationName, opts...)
}

// GlobalTracer returns the global singleton `Tracer` implementation.
// Before `SetGlobalTracer()` is called, the `GlobalTracer()` is a noop
// implementation that drops all data handed to it.
func GlobalTracer() Tracer {
	return opentracing.GlobalTracer()
}
