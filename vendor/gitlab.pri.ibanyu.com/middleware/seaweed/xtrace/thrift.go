package xtrace

import (
	"context"

	"gitlab.pri.ibanyu.com/middleware/seaweed/xcontext"
	"gitlab.pri.ibanyu.com/middleware/seaweed/xtransport/gen-go/util/thriftutil"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

const defaultSpanOpName = "defaultOp"

// TransformOptions keep options for converting a *thriftutil.Context into context.Context
type TransformOptions struct {
	OpName    string
	NoSpan    bool
	NoHead    bool
	NoControl bool
}

// TransformOption defines the interface for any specific option type
type TransformOption interface {
	apply(*TransformOptions)
}

// OpName specify operation name for a generated span, NoSpan can render this option useless.
type OpName string

func (o OpName) apply(tops *TransformOptions) {
	tops.OpName = string(o)
}

// NoSpan stop the conversion from generating new span
type NoSpan bool

func (i NoSpan) apply(topts *TransformOptions) {
	topts.NoSpan = bool(i)
}

// NoHead stop the conversion from keeping the context head
type NoHead bool

func (i NoHead) apply(topts *TransformOptions) {
	topts.NoHead = bool(i)
}

// NoControl stop the conversion from keeping the control head
type NoControl bool

func (i NoControl) apply(topts *TransformOptions) {
	topts.NoControl = bool(i)
}

// NewContextFromThriftUtilContext generate context.Context from *thriftutil.Context, usually this conversion
//   happens when a thrift server receive message from any thrift client.
func NewContextFromThriftUtilContext(tctx *thriftutil.Context, opts ...TransformOption) context.Context {
	topts := &TransformOptions{
		OpName:    defaultSpanOpName,
		NoSpan:    false,
		NoHead:    false,
		NoControl: false,
	}

	for _, opt := range opts {
		opt.apply(topts)
	}
	// defensive
	emptyHead := thriftutil.NewHead()
	emptySpanctx := make(map[string]string)
	emptyControl := thriftutil.NewDefaultControl()
	if tctx == nil {
		tctx = &thriftutil.Context{
			Head:    emptyHead,
			Spanctx: emptySpanctx,
			Control: emptyControl,
		}
	} else {
		if tctx.Head == nil {
			tctx.Head = emptyHead
		}

		if tctx.Spanctx == nil {
			tctx.Spanctx = emptySpanctx
		}

		if tctx.Control == nil {
			tctx.Control = emptyControl
		}
	}

	ctx := context.Background()
	if !topts.NoHead {
		ctx = context.WithValue(ctx, xcontext.ContextKeyHead, tctx.Head)
	}

	if !topts.NoSpan {
		tracer := opentracing.GlobalTracer()
		spanCtx, err := tracer.Extract(opentracing.TextMap, opentracing.TextMapCarrier(tctx.Spanctx))
		var span opentracing.Span
		if err == nil {
			span = tracer.StartSpan(topts.OpName, ext.RPCServerOption(spanCtx))
		} else {
			span = tracer.StartSpan(topts.OpName)
		}
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	if !topts.NoControl {
		ctx = context.WithValue(ctx, xcontext.ContextKeyControl, tctx.Control)
	}

	return ctx
}

// NewThriftUtilContextFromContext generate *thriftutil.Context from context.Context, usually this conversion
//   happens just before triggering a thrift request.
func NewThriftUtilContextFromContext(ctx context.Context, opts ...TransformOption) *thriftutil.Context {
	topts := &TransformOptions{
		NoSpan:    false,
		NoHead:    false,
		NoControl: false,
	}

	for _, opt := range opts {
		opt.apply(topts)
	}

	// defensive
	if ctx == nil {
		ctx = context.Background()
	}

	tctx := &thriftutil.Context{
		Head:    thriftutil.NewHead(),
		Spanctx: make(map[string]string),
		Control: thriftutil.NewDefaultControl(),
	}

	if !topts.NoHead {
		if head, ok := ctx.Value(xcontext.ContextKeyHead).(*thriftutil.Head); ok {
			tctx.Head = head
		}
	}

	if !topts.NoSpan {
		carrier := opentracing.TextMapCarrier(make(map[string]string))
		span := opentracing.SpanFromContext(ctx)
		if span != nil {
			opentracing.GlobalTracer().Inject(
				span.Context(),
				opentracing.TextMap,
				carrier)
		}
		tctx.Spanctx = carrier
	}

	if !topts.NoControl {
		if control, ok := ctx.Value(xcontext.ContextKeyControl).(*thriftutil.Control); ok {
			tctx.Control = control
		}
	}

	return tctx
}
