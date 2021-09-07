package skywalking

import (
	"context"
	"fmt"
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
)

type spanKey struct{}

var (
	tracer          *go2sky.Tracer
	spanKeyInstance = spanKey{}
)

func NewGlobalTracer(serviceName, address string) error {
	r, err := reporter.NewGRPCReporter(address)
	if err != nil {
		return err
	}
	t, err := go2sky.NewTracer(serviceName, go2sky.WithReporter(r))
	if err != nil {
		return err
	}
	tracer = t
	return nil
}

func GetSpanByCtx(ctx context.Context) (Span, error) {
	span, ok := ctx.Value(spanKeyInstance).(Span)
	if ok {
		return span, nil
	}
	return nil, fmt.Errorf("nil in span")
}

func NewRootLocalSpan() (Span, error) {
	rootSpan, ctx, err := tracer.CreateLocalSpan(context.Background())
	if err != nil {
		return nil, err
	}
	span := &skyspan{
		skyWalkingSpan: rootSpan,
	}
	ctx = context.WithValue(ctx, spanKeyInstance, span)
	span.ctx = ctx
	return span, nil
}

func NewRootSpan(operationName string, extractor func(headerKey string) (string, error)) (Span, error) {
	rootSpan, ctx, err := tracer.CreateEntrySpan(context.Background(), operationName, extractor)
	if err != nil {
		return nil, err
	}
	span := &skyspan{
		skyWalkingSpan: rootSpan,
	}
	ctx = context.WithValue(ctx, spanKeyInstance, span)
	span.ctx = ctx
	return span, nil
}
